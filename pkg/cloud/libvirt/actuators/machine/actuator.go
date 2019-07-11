package machine

import (
	"context"
	"fmt"
	"log"

	"sigs.k8s.io/cluster-api-provider-libvirt/pkg/apis/libvirt/v1alpha1"
	l "sigs.k8s.io/cluster-api-provider-libvirt/pkg/cloud/libvirt"
	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

const (
	ProviderName = "libvirt"
)

// Add RBAC rules to access cluster-api resources
//+kubebuilder:rbac:groups=cluster.k8s.io,resources=machines;machines/status,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cluster.k8s.io,resources=machineClasses,verbs=get;list;watch
//+kubebuilder:rbac:groups=cluster.k8s.io,resources=clusters;clusters/status,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=nodes;events,verbs=get;list;watch;create;update;patch;delete

// Actuator is responsible for performing machine reconciliation
type Actuator struct {
	client client.Client
}

func (a *Actuator) Create(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	if machine.Spec.ProviderSpec.Value == nil {
		log.Printf("Machine Provider Spec not passed")
		return fmt.Errorf("Machine Provider Spec not passed")
	}

	var providerSpec v1alpha1.LibvirtMachineProviderSpec
	err := yaml.UnmarshalStrict(machine.Spec.ProviderSpec.Value.Raw, &providerSpec)
	if err != nil {
		log.Printf("Error unmarshalling machine provider spec: %+v", err)
		return err
	}

	spec := providerSpec.Spec
	log.Printf("Create machine actuator called for machine %v", providerSpec)

	err = l.CreateDomain(machine.Name, spec.VCPU, uint(spec.MemoryInGB), spec.ImageURI, spec.UserDataURI)
	if err != nil {
		log.Printf("Error creating node for machine: %v, %v", machine, err)
		return fmt.Errorf("Error creating node for machine: %v, %v", machine, err)
	}

	log.Printf("Machine %v created.", machine.Name)
	return nil
}

func (a *Actuator) Delete(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	panic("not implemented")
}

func (a *Actuator) Update(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	panic("not implemented")
}

func (a *Actuator) Exists(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) (bool, error) {
	return l.DomainExists(machine.Name)
}

func (a *Actuator) GetIP(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (string, error) {
	panic("not implemented")
}

func (a *Actuator) GetKubeConfig(cluster *clusterv1.Cluster, machine *clusterv1.Machine) (string, error) {
	panic("not implemented")
}

type ActuatorParams struct {
	Client client.Client
}

func NewActuator(params ActuatorParams) (*Actuator, error) {
	return &Actuator{
		client: params.Client,
	}, nil
}
