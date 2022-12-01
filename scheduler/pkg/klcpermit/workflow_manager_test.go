package klcpermit

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Test_getCRDName(t *testing.T) {
	Obj := metav1.ObjectMeta{}
	t.Log(len(Obj.Annotations))
	t.Log(Obj.Annotations)
	t.Log(make(map[string]string))

	tests := []struct {
		name string
		pod  *corev1.Pod
		want string
	}{
		{
			name: "properly labeld pod",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						WorkloadAnnotation: "myworkload",
						VersionAnnotation:  "0.0.1",
						AppAnnotation:      "myapp",
					},
				},
			},
			want: "myapp-myworkload-0.0.1",
		},

		{
			name: "properly annotated pod",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						WorkloadAnnotation: "myworkload",
						VersionAnnotation:  "0.0.1",
						AppAnnotation:      "myapp",
					},

					Labels: map[string]string{
						WorkloadAnnotation: "myotherworkload",
						VersionAnnotation:  "0.0.1",
						AppAnnotation:      "mynotapp",
					}},
			},
			want: "myapp-myworkload-0.0.1",
		},

		{
			name: "annotated and labeled pod",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						WorkloadAnnotation: "myworkload",
						VersionAnnotation:  "0.0.1",
						AppAnnotation:      "myapp",
					},
				},
			},
			want: "myapp-myworkload-0.0.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCRDName(tt.pod); got != tt.want {
				t.Errorf("getCRDName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLabelOrAnnotation(t *testing.T) {
	tests := []struct {
		name                string
		pod                 *corev1.Pod
		primaryAnnotation   string
		secondaryAnnotation string
		want                string
		want1               bool
	}{
		{
			name: "Test if primary annotation is returned from annotations",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						AppAnnotation: "some-app-name",
					},
				},
			},
			primaryAnnotation:   AppAnnotation,
			secondaryAnnotation: K8sRecommendedAppAnnotations,
			want:                "some-app-name",
			want1:               true,
		},
		{
			name: "Test if secondary annotation is returned from annotations",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						K8sRecommendedAppAnnotations: "some-app-name",
					},
				},
			},
			primaryAnnotation:   AppAnnotation,
			secondaryAnnotation: K8sRecommendedAppAnnotations,
			want:                "some-app-name",
			want1:               true,
		},
		{
			name: "Test if primary annotation is returned from labels",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						AppAnnotation: "some-app-name",
					},
				},
			},
			primaryAnnotation:   AppAnnotation,
			secondaryAnnotation: K8sRecommendedAppAnnotations,
			want:                "some-app-name",
			want1:               true,
		},
		{
			name: "Test if secondary annotation is returned from labels",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						K8sRecommendedAppAnnotations: "some-app-name",
					},
				},
			},
			primaryAnnotation:   AppAnnotation,
			secondaryAnnotation: K8sRecommendedAppAnnotations,
			want:                "some-app-name",
			want1:               true,
		},
		{
			name: "Test that empty string is returned when no annotations or labels are found",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"some-other-annotation": "some-app-name",
					},
				},
			},
			primaryAnnotation:   AppAnnotation,
			secondaryAnnotation: K8sRecommendedAppAnnotations,
			want:                "",
			want1:               false,
		},
		{
			name: "Test that empty string is returned when primary annotation cannot be found and secondary annotation is empty",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Annotations: map[string]string{
						"some-other-annotation": "some-app-name",
					},
				},
			},
			primaryAnnotation:   AppAnnotation,
			secondaryAnnotation: "",
			want:                "",
			want1:               false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getLabelOrAnnotation(tt.pod, tt.primaryAnnotation, tt.secondaryAnnotation)
			if got != tt.want {
				t.Errorf("getLabelOrAnnotation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getLabelOrAnnotation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getSpan_unbindSpan(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				WorkloadAnnotation: "myworkload",
				VersionAnnotation:  "0.0.1",
				AppAnnotation:      "myapp",
			},
		},
	}

	r := NewWorkloadManager(nil)

	_, span := r.getSpan(context.TODO(), &unstructured.Unstructured{}, pod)

	require.NotNil(t, span)
	require.Len(t, r.bindCRDSpan, 1)

	_, span2 := r.getSpan(context.TODO(), &unstructured.Unstructured{}, pod)

	require.Equal(t, span, span2)
	require.Len(t, r.bindCRDSpan, 1)

	pod2 := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				WorkloadAnnotation: "myworkload",
				VersionAnnotation:  "0.0.2",
				AppAnnotation:      "myapp",
			},
		},
	}

	_, span3 := r.getSpan(context.TODO(), &unstructured.Unstructured{}, pod2)

	require.NotNil(t, span3)
	require.Len(t, r.bindCRDSpan, 2)

	_, span4 := r.getSpan(context.TODO(), &unstructured.Unstructured{}, pod2)

	require.Equal(t, span3, span4)
	require.Len(t, r.bindCRDSpan, 2)

	_, span5 := r.getSpan(context.TODO(), &unstructured.Unstructured{}, pod)

	require.Equal(t, span, span5)
	require.Len(t, r.bindCRDSpan, 2)

	r.unbindSpan(pod)
	require.Len(t, r.bindCRDSpan, 1)

	_, span6 := r.getSpan(context.TODO(), &unstructured.Unstructured{}, pod2)

	require.Equal(t, span3, span6)
	require.Len(t, r.bindCRDSpan, 1)

	_, span7 := r.getSpan(context.TODO(), &unstructured.Unstructured{}, pod)

	require.Equal(t, span, span7)
	require.Len(t, r.bindCRDSpan, 2)

}

func Test_calculateVersion(t *testing.T) {
	tests := []struct {
		name string
		pod  *corev1.Pod
		want string
	}{
		{
			name: "empty pod",
			pod:  &corev1.Pod{},
			want: "2166136261",
		},
		{
			name: "no containers",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "pod-name",
				},
			},
			want: "2166136261",
		},
		{
			name: "single container",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "pod-name",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container-name",
							Image: "image:tag",
						},
					},
				},
			},
			want: "tag",
		},
		{
			name: "single container latest tag",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "pod-name",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container-name",
							Image: "image:latest",
						},
					},
				},
			},
			want: "2894867514",
		},
		{
			name: "multiple containers",
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name: "pod-name",
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "container-name",
							Image: "image:latest",
						},
						{
							Name:  "container-name2",
							Image: "image:latest2",
						},
					},
				},
			},
			want: "3235658121",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateVersion(tt.pod)
			require.Equal(t, tt.want, got)
		})
	}
}
