package docker

import (
	"fmt"

	"github.com/joelferrier/margarita-mixer/config/builder"
)

const (
	defaultApiVersion   = "v1.35"
	defaultUrl          = "unix:///var/run/docker.sock"
	defaultReadTimeout  = 900
	defaultWriteTimeout = 900
)

type DockerBuilder struct {
	url          string
	readTimeout  int64
	writeTimeout int64
	docker       *dockerClient
}

func NewBuilder() *DockerBuilder {
	return &DockerBuilder{
		defaultUrl,          //TODO: pass through to client only
		defaultReadTimeout,  //TODO: pass through to client only
		defaultWriteTimeout, //TODO: pass through to client only
		newClient(defaultUrl),
	}
}

func (b *DockerBuilder) Configure(c builder.Config) {
	fmt.Println("in configure:")
	b.url = c.URL
	b.readTimeout = c.Timeout
	b.writeTimeout = c.Timeout
}

func (b *DockerBuilder) Setup() error {
	fmt.Println("in setup")

	_ = b.docker.images()

	return nil
}

func (b *DockerBuilder) Build() error {
	fmt.Println("in build")
	return nil
}

func (b *DockerBuilder) Extract() error {
	fmt.Println("in extract")
	return nil
}

func (b *DockerBuilder) Cleanup() error {
	fmt.Println("in cleanup")
	return nil
}
