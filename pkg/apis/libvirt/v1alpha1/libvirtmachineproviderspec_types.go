/*
Copyright 2019 GOJEK TECH.

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

// LibvirtMachineProviderSpecSpec defines the desired state of LibvirtMachineProviderSpec
type LibvirtMachineProviderSpecSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Number of virtual CPU
	VCPU int `json:"vcpu"`

	// Amount of RAM in GBs
	MemoryInGB int `json:"memoryInGB"`

	// Image URL to be provisioned
	ImageURI string `json:"imageURI"`

	// UserData URI of cloud-init image
	UserDataURI string `json:"userDataURI"`
}

// LibvirtMachineProviderSpecStatus defines the observed state of LibvirtMachineProviderSpec
type LibvirtMachineProviderSpecStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LibvirtMachineProviderSpec is the Schema for the libvirtmachineproviderspecs API
// +k8s:openapi-gen=true
type LibvirtMachineProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LibvirtMachineProviderSpecSpec   `json:"spec,omitempty"`
	Status LibvirtMachineProviderSpecStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LibvirtMachineProviderSpecList contains a list of LibvirtMachineProviderSpec
type LibvirtMachineProviderSpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LibvirtMachineProviderSpec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LibvirtMachineProviderSpec{}, &LibvirtMachineProviderSpecList{})
}
