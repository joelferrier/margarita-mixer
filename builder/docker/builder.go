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
}

func NewBuilder() *DockerBuilder {
	return &DockerBuilder{
		defaultUrl,
		defaultReadTimeout,
		defaultWriteTimeout,
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
