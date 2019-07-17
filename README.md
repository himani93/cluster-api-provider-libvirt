# cluster-api-provider-libvirt

This repository contains a POC for cluster API libvirt provider. You can use this to create a Kubernetes Control Plane using [libvirt](http://libvirt.org) on a baremetal Linux machine.

## Getting Started

### Prerequisites

Install these prerequisites on a Linux baremetal machine. We have used Ubuntu 18.04 for development and testing.

  * [Golang v1.12.6](https://golang.org/)
  * [Libvirt v4.0.0](#installing-and-configuring-libvirt)
  * [Docker v18.09.7](https://docs.docker.com/install/linux/docker-ce/ubuntu/)
  * [Kubebuilder v1.0.8](https://book-v1.book.kubebuilder.io/getting_started/installation_and_setup.html)
  * [Kustomize v1.0.11](https://github.com/kubernetes-sigs/kustomize/releases/tag/v1.0.11)
  * [Kubectl v1.13](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl-on-linux)
  * [Minikube v1.2.0](https://kubernetes.io/docs/tasks/tools/install-minikube/)

#### Installing and configuring Libvirt

To install Libvirt
```
sudo apt-get install -y \
   qemu-kvm \
   libvirt-daemon-system \
   libvirt-clients \
   bridge-utils \
   libvirt-bin

usermod -a -G libvirt $(whoami)
```
Logout and login for usergroup change to take effect.

Enable TCP listening on Libvirt by editing the `/etc/default/libvirtd` file. It should be configured like this.
```
start_libvirtd="yes"
libvirtd_opts="--listen"
```

Change the `/etc/libvirt/libvirtd.conf` file so that it has the following configuration.

```
# Disable TLS
listen_tls = 0
# Enable TCP port
listen_tcp = 1
# Add TCP port
tcp_port = "16509"
# Setup libvirt socket group
unix_sock_group = "libvirt"
# Setup libvirt socket permissions
unix_sock_ro_perms = "0777"
unix_sock_rw_perms = "0770"
# Setup libvirt auth
auth_unix_ro = "none"
auth_unix_rw = "none"
# Disable TCP auth
auth_tcp = "none"
# Enable auth log
audit_logging = 1
```

Restart libvirtd service.

```
sudo systemctl restart libvirtd
```

### Build and run this provider

  * Clone this repo in `$GOPATH/sigs.k8s.io` directory on the baremetal machine.
    ```
    git clone git@github.com:himani93/cluster-api-provider-libvirt.git $GOPATH/sigs.k8s.io
    ```
  * Run the following commands to build and run a docker image for this project.
    ```
    dep ensure -v
    export IMG=<your docker repo>/cluster-api-provider-libvirt
    make docker-build IMG=${IMG}
    make docker-push IMG=${IMG}
    ```
  * Deploy this provider on a bootstrap minikube cluster.
    ```
    minikube start --vm-driver kvm2
    make deploy
    ```
  * Create a boot disk image as per instructions in [this repo](https://github.com/himani93/vm-builder).
  * Create a file named `user-data`
    ```
    #cloud-config
    password: passw0rd
    chpasswd: {expire: False}
    ssh_pwauth: True
    runcmd:
      - echo "127.0.0.1 kube-cp-01" >> /etc/hosts
      - kubeadm init --pod-network-cidr 10.40.0.0/16
    ```
  * Create a file named `meta-data`
    ```
    instance-id: kube-cp
    local-hostname: kube-cp
    ```
  * Create a cloud init image.
    ```
    genisoimage -output user-data.img -volid cidata -joliet -rock user-data meta-data
    ```
  * Edit the machine CRD file [`machine.yaml`](https://github.com/himani93/cluster-api-provider-libvirt/blob/master/samples/crds/machine.yaml) to make any relevant changes, especially the image paths.
  * Finally run `kubectl apply -f samples/crds/machine.yaml` to create a new machine.

A new virtual machine named `kube-cp` would have been created. You can verify that by running `virsh list`.

**Verify Kuberbetes Control Plane**

Run this command to check the pods that have been created.
```
kubectl --kubeconfig /etc/kubernetes/admin.conf get pods -n kube-system
```



