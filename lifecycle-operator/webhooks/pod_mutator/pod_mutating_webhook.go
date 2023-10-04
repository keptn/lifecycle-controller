package pod_mutator

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	argov1alpha1 "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	"github.com/go-logr/logr"
	klcv1alpha3 "github.com/keptn/lifecycle-toolkit/lifecycle-operator/apis/lifecycle/v1alpha3"
	apicommon "github.com/keptn/lifecycle-toolkit/lifecycle-operator/apis/lifecycle/v1alpha3/common"
	"github.com/keptn/lifecycle-toolkit/lifecycle-operator/apis/lifecycle/v1alpha3/semconv"
	controllercommon "github.com/keptn/lifecycle-toolkit/lifecycle-operator/controllers/common"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/mutate-v1-pod,mutating=true,failurePolicy=fail,groups="",resources=pods,verbs=create;update,versions=v1,name=mpod.keptn.sh,admissionReviewVersions=v1,sideEffects=None
// +kubebuilder:rbac:groups=core,resources=namespaces,verbs=get;list;watch
// +kubebuilder:rbac:groups=apps,resources=deployments;statefulsets;daemonsets;replicasets,verbs=get

// PodMutatingWebhook annotates Pods
type PodMutatingWebhook struct {
	Client                 client.Client
	Tracer                 trace.Tracer
	Decoder                *admission.Decoder
	EventSender            controllercommon.IEvent
	Log                    logr.Logger
	SchedulingGatesEnabled bool
}

const InvalidAnnotationMessage = "Invalid annotations"

var ErrTooLongAnnotations = fmt.Errorf("too long annotations, maximum length for app and workload is 25 characters, for version 12 characters")

// Handle inspects incoming Pods and injects the Keptn scheduler if they contain the Keptn lifecycle annotations.
//
//nolint:gocyclo
func (a *PodMutatingWebhook) Handle(ctx context.Context, req admission.Request) admission.Response {

	ctx, span := a.Tracer.Start(ctx, "annotate_pod", trace.WithNewRoot(), trace.WithSpanKind(trace.SpanKindServer))
	defer span.End()

	a.Log.Info("webhook for pod called")

	pod := &corev1.Pod{}

	err := a.Decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	// check if Lifecycle Operator is enabled for this namespace
	namespace := &corev1.Namespace{}
	if err = a.Client.Get(ctx, types.NamespacedName{Name: req.Namespace}, namespace); err != nil {
		a.Log.Error(err, "could not get namespace", "namespace", req.Namespace)
		return admission.Errored(http.StatusInternalServerError, err)
	}

	if namespace.GetAnnotations()[apicommon.NamespaceEnabledAnnotation] != "enabled" {
		a.Log.Info("namespace is not enabled for lifecycle operator", "namespace", req.Namespace)
		return admission.Allowed("namespace is not enabled for lifecycle operator")
	}

	// check the OwnerReference of the pod to see if it is supported and intended to be managed by Keptn
	ownerRef := getOwnerReference(&pod.ObjectMeta)

	if ownerRef.Kind == "" {
		msg := "owner of pod is not supported by lifecycle operator"
		a.Log.Info(msg, "namespace", req.Namespace, "pod", req.Name)
		return admission.Allowed(msg)
	}

	a.Log.Info(fmt.Sprintf("Pod annotations: %v", pod.Annotations))

	podIsAnnotated := isPodAnnotated(pod)
	if !podIsAnnotated {
		a.Log.Info("Pod is not annotated, check for parent annotations...")
		podIsAnnotated = a.copyAnnotationsIfParentAnnotated(ctx, &req, pod)
	}

	if podIsAnnotated {
		a.Log.Info("Resource is annotated with Keptn annotations")

		if scheduled := handleScheduling(a, a.Log, pod); scheduled {
			return admission.Allowed("gate of the pod already removed")
		}

		a.Log.Info("Annotations", "annotations", pod.Annotations)
		semconv.AddAttributeFromAnnotations(span, pod.Annotations)
		a.Log.Info("Attributes from annotations set")

		if err := a.handleWorkload(ctx, pod, req.Namespace); err != nil {
			a.Log.Error(err, "Could not handle Workload")
			span.SetStatus(codes.Error, err.Error())
			return admission.Errored(http.StatusBadRequest, err)
		}

		if err := a.handleApp(ctx, pod, req.Namespace); err != nil {
			a.Log.Error(err, "Could not handle App")
			span.SetStatus(codes.Error, err.Error())
			return admission.Errored(http.StatusBadRequest, err)
		}
	}

	marshaledPod, err := json.Marshal(pod)
	if err != nil {
		span.SetStatus(codes.Error, "Failed to marshal")
		return admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}

func handleScheduling(a *PodMutatingWebhook, logger logr.Logger, pod *corev1.Pod) bool {
	if a.SchedulingGatesEnabled {
		logger.Info("SchedulingGates enabled")
		_, gateRemoved := getLabelOrAnnotation(&pod.ObjectMeta, apicommon.SchedulingGateRemoved, "")
		if gateRemoved {
			return true
		}
		pod.Spec.SchedulingGates = []corev1.PodSchedulingGate{
			{
				Name: apicommon.KeptnGate,
			},
		}
	} else {
		a.Log.Info("SchedulingGates disabled, using keptn-scheduler")
		pod.Spec.SchedulerName = "keptn-scheduler"
	}
	return false
}

func (a *PodMutatingWebhook) copyAnnotationsIfParentAnnotated(ctx context.Context, req *admission.Request, pod *corev1.Pod) bool {
	podOwner := getOwnerReference(&pod.ObjectMeta)
	if podOwner.UID == "" {
		return false
	}

	switch podOwner.Kind {
	case "ReplicaSet":
		rs := &appsv1.ReplicaSet{}
		if err := a.Client.Get(ctx, types.NamespacedName{Namespace: req.Namespace, Name: podOwner.Name}, rs); err != nil {
			return false
		}
		a.Log.Info("Done fetching RS")

		rsOwner := getOwnerReference(&rs.ObjectMeta)
		if rsOwner.UID == "" {
			return false
		}

		if rsOwner.Kind == "Rollout" {
			ro := &argov1alpha1.Rollout{}
			return a.fetchParentObjectAndCopyLabels(ctx, podOwner.Name, req.Namespace, pod, ro)
		}
		dp := &appsv1.Deployment{}
		return a.fetchParentObjectAndCopyLabels(ctx, rsOwner.Name, req.Namespace, pod, dp)

	case "StatefulSet":
		sts := &appsv1.StatefulSet{}
		return a.fetchParentObjectAndCopyLabels(ctx, podOwner.Name, req.Namespace, pod, sts)
	case "DaemonSet":
		ds := &appsv1.DaemonSet{}
		return a.fetchParentObjectAndCopyLabels(ctx, podOwner.Name, req.Namespace, pod, ds)
	default:
		return false
	}
}

func (a *PodMutatingWebhook) fetchParentObjectAndCopyLabels(ctx context.Context, name string, namespace string, pod *corev1.Pod, objectContainer client.Object) bool {
	if err := a.Client.Get(ctx, types.NamespacedName{Namespace: namespace, Name: name}, objectContainer); err != nil {
		return false
	}
	objectContainerMetaData := metav1.ObjectMeta{
		Labels:      objectContainer.GetLabels(),
		Annotations: objectContainer.GetAnnotations(),
	}
	return copyResourceLabelsIfPresent(&objectContainerMetaData, pod)
}

func copyResourceLabelsIfPresent(sourceResource *metav1.ObjectMeta, targetPod *corev1.Pod) bool {
	var workloadName, appName, version, preDeploymentChecks, postDeploymentChecks, preEvaluationChecks, postEvaluationChecks string
	var gotWorkloadName, gotVersion bool

	workloadName, gotWorkloadName = getLabelOrAnnotation(sourceResource, apicommon.WorkloadAnnotation, apicommon.K8sRecommendedWorkloadAnnotations)
	appName, _ = getLabelOrAnnotation(sourceResource, apicommon.AppAnnotation, apicommon.K8sRecommendedAppAnnotations)
	version, gotVersion = getLabelOrAnnotation(sourceResource, apicommon.VersionAnnotation, apicommon.K8sRecommendedVersionAnnotations)
	preDeploymentChecks, _ = getLabelOrAnnotation(sourceResource, apicommon.PreDeploymentTaskAnnotation, "")
	postDeploymentChecks, _ = getLabelOrAnnotation(sourceResource, apicommon.PostDeploymentTaskAnnotation, "")
	preEvaluationChecks, _ = getLabelOrAnnotation(sourceResource, apicommon.PreDeploymentEvaluationAnnotation, "")
	postEvaluationChecks, _ = getLabelOrAnnotation(sourceResource, apicommon.PostDeploymentEvaluationAnnotation, "")

	if len(targetPod.Annotations) == 0 {
		targetPod.Annotations = make(map[string]string)
	}

	if gotWorkloadName {
		setMapKey(targetPod.Annotations, apicommon.WorkloadAnnotation, workloadName)

		if !gotVersion {
			setMapKey(targetPod.Annotations, apicommon.VersionAnnotation, calculateVersion(targetPod))
		} else {
			setMapKey(targetPod.Annotations, apicommon.VersionAnnotation, version)
		}

		setMapKey(targetPod.Annotations, apicommon.AppAnnotation, appName)
		setMapKey(targetPod.Annotations, apicommon.PreDeploymentTaskAnnotation, preDeploymentChecks)
		setMapKey(targetPod.Annotations, apicommon.PostDeploymentTaskAnnotation, postDeploymentChecks)
		setMapKey(targetPod.Annotations, apicommon.PreDeploymentEvaluationAnnotation, preEvaluationChecks)
		setMapKey(targetPod.Annotations, apicommon.PostDeploymentEvaluationAnnotation, postEvaluationChecks)

		return true
	}
	return false
}

//nolint:dupl
func (a *PodMutatingWebhook) handleWorkload(ctx context.Context, pod *corev1.Pod, namespace string) error {

	ctx, span := a.Tracer.Start(ctx, "create_workload", trace.WithSpanKind(trace.SpanKindProducer))
	defer span.End()

	newWorkload := generateWorkload(ctx, pod, namespace)

	newWorkload.SetSpanAttributes(span)

	a.Log.Info("Searching for workload")

	workload := &klcv1alpha3.KeptnWorkload{}
	err := a.Client.Get(ctx, types.NamespacedName{Namespace: namespace, Name: newWorkload.Name}, workload)
	if errors.IsNotFound(err) {
		a.Log.Info("Creating workload", "workload", workload.Name)
		workload = newWorkload
		err = a.Client.Create(ctx, workload)
		if err != nil {
			a.Log.Error(err, "Could not create Workload")
			a.EventSender.Emit(apicommon.PhaseCreateWorkload, "Warning", workload, apicommon.PhaseStateFailed, "could not create KeptnWorkload", workload.Spec.Version)
			span.SetStatus(codes.Error, err.Error())
			return err
		}

		return nil
	}

	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return fmt.Errorf("could not fetch Workload"+": %+v", err)
	}

	if reflect.DeepEqual(workload.Spec, newWorkload.Spec) {
		a.Log.Info("Pod not changed, not updating anything")
		return nil
	}

	a.Log.Info("Pod changed, updating workload")
	workload.Spec = newWorkload.Spec

	err = a.Client.Update(ctx, workload)
	if err != nil {
		a.Log.Error(err, "Could not update Workload")
		a.EventSender.Emit(apicommon.PhaseUpdateWorkload, "Warning", workload, apicommon.PhaseStateFailed, "could not update KeptnWorkload", workload.Spec.Version)
		span.SetStatus(codes.Error, err.Error())
		return err
	}

	return nil
}

//nolint:dupl
func (a *PodMutatingWebhook) handleApp(ctx context.Context, pod *corev1.Pod, namespace string) error {

	ctx, span := a.Tracer.Start(ctx, "create_app", trace.WithSpanKind(trace.SpanKindProducer))
	defer span.End()

	newAppCreationRequest := generateAppCreationRequest(ctx, pod, namespace)

	newAppCreationRequest.SetSpanAttributes(span)

	a.Log.Info("Searching for AppCreationRequest", "appCreationRequest", newAppCreationRequest.Name)

	appCreationRequest := &klcv1alpha3.KeptnAppCreationRequest{}
	err := a.Client.Get(ctx, types.NamespacedName{Namespace: namespace, Name: newAppCreationRequest.Name}, appCreationRequest)
	if errors.IsNotFound(err) {
		a.Log.Info("Creating app creation request", "appCreationRequest", appCreationRequest.Name)
		appCreationRequest = newAppCreationRequest
		err = a.Client.Create(ctx, appCreationRequest)
		if err != nil {
			a.Log.Error(err, "Could not create AppCreationRequest")
			a.EventSender.Emit(apicommon.PhaseCreateAppCreationRequest, "Warning", appCreationRequest, apicommon.PhaseStateFailed, "could not create KeptnAppCreationRequest", appCreationRequest.Spec.AppName)
			span.SetStatus(codes.Error, err.Error())
			return err
		}

		return nil
	}

	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		return fmt.Errorf("could not fetch AppCreationRequest"+": %+v", err)
	}

	return nil
}

func generateWorkload(ctx context.Context, pod *corev1.Pod, namespace string) *klcv1alpha3.KeptnWorkload {
	version := getVersion(&pod.ObjectMeta)
	applicationName := getAppName(&pod.ObjectMeta)

	var preDeploymentTasks []string
	var postDeploymentTasks []string
	var preDeploymentEvaluation []string
	var postDeploymentEvaluation []string

	if annotations, found := getLabelOrAnnotation(&pod.ObjectMeta, apicommon.PreDeploymentTaskAnnotation, ""); found {
		preDeploymentTasks = strings.Split(annotations, ",")
	}

	if annotations, found := getLabelOrAnnotation(&pod.ObjectMeta, apicommon.PostDeploymentTaskAnnotation, ""); found {
		postDeploymentTasks = strings.Split(annotations, ",")
	}

	if annotations, found := getLabelOrAnnotation(&pod.ObjectMeta, apicommon.PreDeploymentEvaluationAnnotation, ""); found {
		preDeploymentEvaluation = strings.Split(annotations, ",")
	}

	if annotations, found := getLabelOrAnnotation(&pod.ObjectMeta, apicommon.PostDeploymentEvaluationAnnotation, ""); found {
		postDeploymentEvaluation = strings.Split(annotations, ",")
	}

	// create TraceContext
	// follow up with a Keptn propagator that JSON-encoded the OTel map into our own key
	traceContextCarrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, traceContextCarrier)

	ownerRef := getOwnerReference(&pod.ObjectMeta)

	return &klcv1alpha3.KeptnWorkload{
		ObjectMeta: metav1.ObjectMeta{
			Name:        getWorkloadName(&pod.ObjectMeta),
			Namespace:   namespace,
			Annotations: traceContextCarrier,
			OwnerReferences: []metav1.OwnerReference{
				ownerRef,
			},
		},
		Spec: klcv1alpha3.KeptnWorkloadSpec{
			AppName:                   applicationName,
			Version:                   version,
			ResourceReference:         klcv1alpha3.ResourceReference{UID: ownerRef.UID, Kind: ownerRef.Kind, Name: ownerRef.Name},
			PreDeploymentTasks:        preDeploymentTasks,
			PostDeploymentTasks:       postDeploymentTasks,
			PreDeploymentEvaluations:  preDeploymentEvaluation,
			PostDeploymentEvaluations: postDeploymentEvaluation,
		},
	}
}

func generateAppCreationRequest(ctx context.Context, pod *corev1.Pod, namespace string) *klcv1alpha3.KeptnAppCreationRequest {

	appName := getAppName(&pod.ObjectMeta)

	// create TraceContext
	// follow up with a Keptn propagator that JSON-encoded the OTel map into our own key
	traceContextCarrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, traceContextCarrier)

	kacr := &klcv1alpha3.KeptnAppCreationRequest{
		ObjectMeta: metav1.ObjectMeta{
			Name:        appName,
			Namespace:   namespace,
			Annotations: traceContextCarrier,
		},
		Spec: klcv1alpha3.KeptnAppCreationRequestSpec{
			AppName: appName,
		},
	}

	if !isAppAnnotationPresent(pod) {
		kacr.Annotations[apicommon.AppTypeAnnotation] = string(apicommon.AppTypeSingleService)
	}

	return kacr
}
