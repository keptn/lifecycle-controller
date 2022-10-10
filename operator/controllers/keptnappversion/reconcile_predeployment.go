package keptnappversion

import (
	"context"

	klcv1alpha1 "github.com/keptn-sandbox/lifecycle-controller/operator/api/v1alpha1"
	"github.com/keptn-sandbox/lifecycle-controller/operator/api/v1alpha1/common"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *KeptnAppVersionReconciler) reconcilePreDeployment(ctx context.Context, req ctrl.Request, appVersion *klcv1alpha1.KeptnAppVersion) (common.KeptnState, error) {
	newStatus, preDeploymentState, err := r.reconcileChecks(ctx, common.PreDeploymentCheckType, appVersion)
	if err != nil {
		return common.StateUnknown, err
	}
	overallState := common.GetOverallState(preDeploymentState)
	appVersion.Status.PreDeploymentStatus = overallState
	appVersion.Status.PreDeploymentTaskStatus = newStatus

	// Write Status Field
	err = r.Client.Status().Update(ctx, appVersion)
	if err != nil {
		return common.StateUnknown, err
	}
	return overallState, nil
}
