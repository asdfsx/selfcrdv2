/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SelfCRDV2Spec defines the desired state of SelfCRDV2
type SelfCRDV2Spec struct {
	// +kubebuilder:validation:MinLength=0
	Username string `json:"username,omitempty"`

	// +kubebuilder:validation:MinLength=0
	CustomID string `json:"custom_id,omitempty"`
}

// SelfCRDV2Status defines the observed state of SelfCRDV2
type SelfCRDV2Status struct {

	// A list of pointers to currently running jobs.
	// +optional
	Active []corev1.ObjectReference `json:"active,omitempty"`

	// Information when was the last time the job was successfully scheduled.
	// +optional
	LastScheduleTime *metav1.Time `json:"lastScheduleTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// SelfCRDV2 is the Schema for the selfcrdv2s API
type SelfCRDV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SelfCRDV2Spec   `json:"spec,omitempty"`
	Status SelfCRDV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SelfCRDV2List contains a list of SelfCRDV2
type SelfCRDV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SelfCRDV2 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SelfCRDV2{}, &SelfCRDV2List{})
}
