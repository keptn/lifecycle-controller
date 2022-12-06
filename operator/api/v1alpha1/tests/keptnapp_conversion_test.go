package api

import (
	"github.com/keptn/lifecycle-toolkit/operator/api/v1alpha1"
	"github.com/keptn/lifecycle-toolkit/operator/api/v1alpha2"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestKeptnApp_ConvertFrom(t *testing.T) {
	tests := []struct {
		name    string
		srcObj  *v1alpha2.KeptnApp
		wantErr bool
		wantObj *v1alpha1.KeptnApp
	}{
		{
			name: "Test that conversion from v1alpha2 to v1alpha1 works",
			srcObj: &v1alpha2.KeptnApp{
				TypeMeta: v1.TypeMeta{
					Kind:       "KeptnApp",
					APIVersion: "lifecycle.keptn.sh/v1alpha2",
				},
				ObjectMeta: v1.ObjectMeta{
					Name:      "some-keptn-app-name",
					Namespace: "",
					Labels: map[string]string{
						"some-label": "some-label-value",
					},
					Annotations: map[string]string{
						"some-annotation": "some-annotation-value",
					},
				},
				Spec: v1alpha2.KeptnAppSpec{
					Version:  "1.2.3",
					Revision: "1",
					Workloads: []v1alpha2.KeptnWorkloadRef{
						{
							Name:    "workload-1",
							Version: "1.2.3",
						},
						{
							Name:    "workload-2",
							Version: "4.5.6",
						},
					},
					PreDeploymentTasks: []string{
						"some-pre-deployment-task1",
					},
					PostDeploymentTasks: []string{
						"some-post-deployment-task2",
					},
					PreDeploymentEvaluations: []string{
						"some-pre-evaluation-task1",
					},
					PostDeploymentEvaluations: []string{
						"some-pre-evaluation-task2",
					},
				},
				Status: v1alpha2.KeptnAppStatus{
					CurrentVersion: "1.2.3",
				},
			},
			wantErr: false,
			wantObj: &v1alpha1.KeptnApp{
				TypeMeta: v1.TypeMeta{
					Kind:       "KeptnApp",
					APIVersion: "lifecycle.keptn.sh/v1alpha2",
				},
				ObjectMeta: v1.ObjectMeta{
					Name:      "some-keptn-app-name",
					Namespace: "",
					Labels: map[string]string{
						"some-label": "some-label-value",
					},
					Annotations: map[string]string{
						"some-annotation": "some-annotation-value",
					},
				},
				Spec: v1alpha1.KeptnAppSpec{
					Version: "1.2.3",
					Workloads: []v1alpha1.KeptnWorkloadRef{
						{
							Name:    "workload-1",
							Version: "1.2.3",
						},
						{
							Name:    "workload-2",
							Version: "4.5.6",
						},
					},
					PreDeploymentTasks: []string{
						"some-pre-deployment-task1",
					},
					PostDeploymentTasks: []string{
						"some-post-deployment-task2",
					},
					PreDeploymentEvaluations: []string{
						"some-pre-evaluation-task1",
					},
					PostDeploymentEvaluations: []string{
						"some-pre-evaluation-task2",
					},
				},
				Status: v1alpha1.KeptnAppStatus{
					CurrentVersion: "1.2.3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := &v1alpha1.KeptnApp{
				TypeMeta:   v1.TypeMeta{},
				ObjectMeta: v1.ObjectMeta{},
				Spec:       v1alpha1.KeptnAppSpec{},
				Status:     v1alpha1.KeptnAppStatus{},
			}
			if err := dst.ConvertFrom(tt.srcObj); (err != nil) != tt.wantErr {
				t.Errorf("ConvertFrom() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantObj != nil {
				require.Equal(t, tt.wantObj, dst, "Object was not converted correctly")
			}
		})
	}
}

func TestKeptnApp_ConvertTo(t *testing.T) {
	tests := []struct {
		name    string
		src     *v1alpha1.KeptnApp
		wantErr bool
		wantObj *v1alpha2.KeptnApp
	}{
		{
			name: "Test that conversion from v1alpha1 to v1alpha2 works",
			src: &v1alpha1.KeptnApp{
				TypeMeta: v1.TypeMeta{
					Kind:       "KeptnApp",
					APIVersion: "lifecycle.keptn.sh/v1alpha2",
				},
				ObjectMeta: v1.ObjectMeta{
					Name:      "some-keptn-app-name",
					Namespace: "",
					Labels: map[string]string{
						"some-label": "some-label-value",
					},
					Annotations: map[string]string{
						"some-annotation": "some-annotation-value",
					},
				},
				Spec: v1alpha1.KeptnAppSpec{
					Version: "1.2.3",
					Workloads: []v1alpha1.KeptnWorkloadRef{
						{
							Name:    "workload-1",
							Version: "1.2.3",
						},
						{
							Name:    "workload-2",
							Version: "4.5.6",
						},
					},
					PreDeploymentTasks: []string{
						"some-pre-deployment-task1",
					},
					PostDeploymentTasks: []string{
						"some-post-deployment-task2",
					},
					PreDeploymentEvaluations: []string{
						"some-pre-evaluation-task1",
					},
					PostDeploymentEvaluations: []string{
						"some-pre-evaluation-task2",
					},
				},
				Status: v1alpha1.KeptnAppStatus{
					CurrentVersion: "1.2.3",
				},
			},
			wantErr: false,
			wantObj: &v1alpha2.KeptnApp{
				TypeMeta: v1.TypeMeta{
					Kind:       "KeptnApp",
					APIVersion: "lifecycle.keptn.sh/v1alpha2",
				},
				ObjectMeta: v1.ObjectMeta{
					Name:      "some-keptn-app-name",
					Namespace: "",
					Labels: map[string]string{
						"some-label": "some-label-value",
					},
					Annotations: map[string]string{
						"some-annotation": "some-annotation-value",
					},
				},
				Spec: v1alpha2.KeptnAppSpec{
					Version:  "1.2.3",
					Revision: "1",
					Workloads: []v1alpha2.KeptnWorkloadRef{
						{
							Name:    "workload-1",
							Version: "1.2.3",
						},
						{
							Name:    "workload-2",
							Version: "4.5.6",
						},
					},
					PreDeploymentTasks: []string{
						"some-pre-deployment-task1",
					},
					PostDeploymentTasks: []string{
						"some-post-deployment-task2",
					},
					PreDeploymentEvaluations: []string{
						"some-pre-evaluation-task1",
					},
					PostDeploymentEvaluations: []string{
						"some-pre-evaluation-task2",
					},
				},
				Status: v1alpha2.KeptnAppStatus{
					CurrentVersion: "1.2.3",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dst := v1alpha2.KeptnApp{
				TypeMeta:   v1.TypeMeta{},
				ObjectMeta: v1.ObjectMeta{},
				Spec:       v1alpha2.KeptnAppSpec{},
				Status:     v1alpha2.KeptnAppStatus{},
			}
			if err := tt.src.ConvertTo(&dst); (err != nil) != tt.wantErr {
				t.Errorf("ConvertTo() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantObj != nil {
				require.Equal(t, tt.wantObj, &dst, "Object was not converted correctly")
			}
		})
	}
}
