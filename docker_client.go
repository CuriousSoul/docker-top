package main

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

//const socketPath string = "unix://var/run/docker.sock"
const socketPath string = "/var/run/com.docker.vmnetd.sock"

// DockerClient contains client data structure
type DockerClient struct {
	conn *client.Client
}

// DockerInterface defines "fat" docker interface -- TODO: Refactor this
type DockerInterface interface {
	ListContainers() []types.Container
	ListNetworks() []types.NetworkResource
	ListVolumes() []*types.Volume // Darn,why is this pointer when everything else is not?
}

// NewDockerClient returns docker client
func NewDockerClient() (*DockerClient, error) {
	clnt, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &DockerClient{conn: clnt}, nil
}

// ListContainers provides the list of running containers
func (clnt *DockerClient) ListContainers() []types.Container {
	containers, err := clnt.conn.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic("Error in call ContainerList")
	}
	return containers
}

// ListNetworks provides the list of docker networks
func (clnt *DockerClient) ListNetworks() []types.NetworkResource {
	networks, err := clnt.conn.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic("Error in call NetworkList")
	}
	return networks
}

// ListVolumes provides the list of docker volumes
func (clnt *DockerClient) ListVolumes() []*types.Volume {
	volumes, err := clnt.conn.VolumeList(context.Background(), filters.Args{})
	if err != nil {
		panic("Error in call VolumeList")
	}
	return volumes.Volumes
}
