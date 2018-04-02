package docker

import (
	"fmt"
	"time"

	"github.com/joelferrier/margarita-mixer/config/builder"
	"github.com/joelferrier/margarita-mixer/config/profile"
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
	containers   map[string]*container
}

func NewBuilder() (*DockerBuilder, error) {
	client, err := newClient(
		defaultUrl,
		defaultApiVersion,
		defaultReadTimeout,
		defaultWriteTimeout,
	)

	containers := make(map[string]*container)

	return &DockerBuilder{
		defaultUrl,
		defaultReadTimeout,
		defaultWriteTimeout,
		client,
		containers,
	}, err
}

func (b *DockerBuilder) Configure(c builder.Config) {
	fmt.Println("in configure:")
	b.url = c.URL
	b.readTimeout = c.Timeout
	b.writeTimeout = c.Timeout
}

func (b *DockerBuilder) Setup(p profile.Config) error {
	fmt.Println("in setup")
	var err error
	_ = b.docker.images()
	err = b.docker.pull(p.Image)
	if err != nil {
		return err
	}
	_ = b.docker.images()
	c, err := b.docker.container(p.Image)
	if err != nil {
		return err
	}
	b.containers[p.Image] = c

	fmt.Println(c)
	_ = c.start()

	return err
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
	var err error

	for name, container := range b.containers {
		time.Sleep(100)
		fmt.Printf("stopping %s container\n", name)
		err = container.stop()
		if err != nil {
			return err
		}
		time.Sleep(30)
		fmt.Printf("removing %s container\n", name)
		_ = container.remove()
		if err != nil {
			return err
		}
	}
	return err //TODO: return slice of errors
}
