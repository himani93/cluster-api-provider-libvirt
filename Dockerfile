# Build the manager binary
FROM golang:1.10.3 as builder

# Copy in the go src
WORKDIR /go/src/sigs.k8s.io/cluster-api-provider-libvirt
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/

# Build
RUN apt-get update && apt-get install -y libvirt-dev
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -a -o manager sigs.k8s.io/cluster-api-provider-libvirt/cmd/manager

# Copy the controller-manager into a thin image
FROM ubuntu:latest
WORKDIR /
COPY --from=builder /go/src/sigs.k8s.io/cluster-api-provider-libvirt/manager .
ENTRYPOINT ["/manager"]
