/*
Copyright 2025.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// O2ImsSpec defines the desired state of O2Ims
type O2ImsSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Endpoint defines the o2ims endpoint to call
	Endpoint string `json:"endpoint,omitempty"`
}

// O2ImsStatus defines the observed state of O2Ims
type O2ImsStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// O2Ims is the Schema for the o2ims API
type O2Ims struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   O2ImsSpec   `json:"spec,omitempty"`
	Status O2ImsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// O2ImsList contains a list of O2Ims
type O2ImsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []O2Ims `json:"items"`
}

func init() {
	SchemeBuilder.Register(&O2Ims{}, &O2ImsList{})
}
