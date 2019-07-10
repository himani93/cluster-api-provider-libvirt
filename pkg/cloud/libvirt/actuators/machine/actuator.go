package machine

import (
	"context"

	clusterv1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	ProviderName = "libvirt"
)

// Actuator is responsible for performing machine reconciliation
type Actuator struct {
	client client.Client
}

func (a *Actuator) Create(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	panic("not implemented")
}

func (a *Actuator) Delete(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	panic("not implemented")
}

func (a *Actuator) Update(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) error {
	panic("not implemented")
}

func (a *Actuator) Exists(ctx context.Context, cluster *clusterv1.Cluster, machine *clusterv1.Machine) (bool, error) {
	panic("not implemented")
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
