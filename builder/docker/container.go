package docker

import (
	"github.com/docker/docker/api/types"
	container_t "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

var parkCommand = []string{"bash", "-c", "/usr/bin/tail -f /dev/null"}
var startOptions = types.ContainerStartOptions{}
var removeOptions = types.ContainerRemoveOptions{}

type container struct {
	ID   string
	Name string
	cli  *client.Client
}

func newContainer(name string, image string, cli *client.Client) (*container, error) {
	ctx := context.Background()
	resp, err := cli.ContainerCreate(ctx, &container_t.Config{
		Image: image,
		Cmd:   parkCommand,
	}, nil, nil, name)
	return &container{
		resp.ID,
		name,
		cli,
	}, err
}

func (c *container) start() error {
	ctx := context.Background()
	return c.cli.ContainerStart(ctx, c.ID, startOptions)
}

func (c *container) stop() error {
	ctx := context.Background()
	return c.cli.ContainerStop(ctx, c.ID, nil)
}

func (c *container) remove() error {
	ctx := context.Background()
	return c.cli.ContainerRemove(ctx, c.ID, removeOptions)
}
