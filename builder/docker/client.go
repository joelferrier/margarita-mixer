package docker

import (
	"fmt"

	"github.com/docker/docker/api/types"
	//"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type dockerClient struct {
	url string
	cli *client.Client
}

func newClient(url string) *dockerClient {
	fmt.Println("in newClient")
	return &dockerClient{
		url,
		nil,
	}
}

func (c *dockerClient) images() error {
	fmt.Println("in images")
	var err error

	c.cli, err = client.NewClient(c.url, defaultApiVersion, nil, nil)
	if err != nil {
		panic(err)
	}

	images, err := c.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}

	return nil
}
