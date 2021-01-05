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

package v1alpha1

import (
	"gitlab-ce.alauda.cn/asm/flagger-operator/pkg/oprlib/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FlaggerSpec defines the desired state of Flagger
type FlaggerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Flagger. Edit Flagger_types.go to remove/update
	//Foo string `json:"foo,omitempty"`
	Parameters string `json:"parameters,omitempty"`
}

// FlaggerStatus defines the observed state of Flagger
type FlaggerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	api.OperatorStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Flagger is the Schema for the flaggers API
type Flagger struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlaggerSpec   `json:"spec,omitempty"`
	Status FlaggerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlaggerList contains a list of Flagger
type FlaggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Flagger `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Flagger{}, &FlaggerList{})
}

func (in Flagger) GetOperatorParams() (map[string]interface{}, error) {
	var m map[string]interface{}
	if err := yaml.Unmarshal([]byte(in.Spec.Parameters), &m); err != nil {
		return nil, err
	}

	return m, nil
}
