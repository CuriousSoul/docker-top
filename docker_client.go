package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

const socketPath string = "unix://var/run/docker.sock"

// DockerClient contains client data structure
type DockerClient struct {
	conn *client.Client
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
func (clnt *DockerClient) ListContainers() {
	containers, err := clnt.conn.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic("Error")
	}
	for _, container := range containers {
		fmt.Printf("%s %s %s\n", container.Names[0], container.State, container.Image)
	}
}

// ListNetworks provides the list of docker networks
func (clnt *DockerClient) ListNetworks() {
	networks, err := clnt.conn.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		panic("Error")
	}
	for _, network := range networks {
		fmt.Printf("%s %s\n", network.Name, network.ID)
	}
}

// ListVolumes provides the list of docker volumes
func (clnt *DockerClient) ListVolumes() {
	volumes, err := clnt.conn.VolumeList(context.Background(), filters.Args{})
	if err != nil {
		panic("Error")
	}

	for _, volume := range volumes.Volumes {
		fmt.Printf("%s %s\n", volume.Name, volume.Driver)
	}
}
