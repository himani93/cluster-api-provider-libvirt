package controller

import (
	"sigs.k8s.io/cluster-api-provider-libvirt/pkg/cloud/libvirt/actuators/machine"
	capimachine "sigs.k8s.io/cluster-api/pkg/controller/machine"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

//+kubebuilder:rbac:groups=libvirt.cluster.k8s.io,resources=hwmachineproviderspecs;hwmachineproviderstatuses,verbs=get;list;watch;create;update;patch;delete
func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, func(m manager.Manager) error {
		return capimachine.AddWithActuator(m, &machine.Actuator{})
	})
}
