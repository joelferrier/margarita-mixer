package docker

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
	"golang.org/x/net/context"
)

type dockerClient struct {
	url          string
	apiVersion   string
	readTimeout  int64
	writeTimeout int64
	cli          *client.Client
}

func newClient(url string, apiVersion string, readTimeout int64, writeTimeout int64) (*dockerClient, error) {
	cli, err := client.NewClient(url, apiVersion, nil, nil)

	return &dockerClient{
		url,
		apiVersion,
		readTimeout,
		writeTimeout,
		cli,
	}, err
}

func (c *dockerClient) images() error {
	fmt.Println("in images")
	var err error

	images, err := c.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		var t = "untagged"
		if len(image.RepoTags) > 0 {
			t = image.RepoTags[0]
		}

		fmt.Printf("%s ~> %s\n", t, image.ID[7:19])
	}

	return nil
}

func (c *dockerClient) pull(image string) error {
	var err error
	ctx := context.Background()

	events, err := c.cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer events.Close()
	d := json.NewDecoder(events)

	var event *pullEvent
	var bar *mpb.Bar
	var lastProgress = 0

	p := mpb.New(
		mpb.WithWidth(100),
	)

	for {
		err := d.Decode(&event)

		if err == io.EOF {
			p.Stop()
			break
		} else if err != nil {
			p.Stop()
			return err
		}

		switch status := event.Status; status {
		case "Downloading":
			fallthrough
		case "Extracting":
			if bar == nil || bar.InProgress() == false {
				lastProgress = 0
				bar = p.AddBar(int64(event.ProgressDetail.Total),
					mpb.PrependDecorators(
						decor.StaticName(status, 11, 0),
						decor.ETA(4, 0),
					),
					mpb.AppendDecorators(
						decor.Percentage(5, 0),
					),
				)
			} else {
				bar.IncrBy(event.ProgressDetail.Current - lastProgress)
				lastProgress = event.ProgressDetail.Current
			}
		case "Download complete":
			fallthrough
		case "Verifying checksum":
			fallthrough
		case "Pull complete":
			if bar != nil || bar.InProgress() {
				bar.SetTotal(bar.Total(), true)
				bar.Complete()
				p.RemoveBar(bar)
			}
			fmt.Println(status)
		default:
			fmt.Println(status)
		}
	}

	return err
}

func (c *dockerClient) container(image string) (*container, error) {
	var name string
	parts := strings.Split(image, ":")
	name = fmt.Sprintf("mixer_%s", parts[0])
	return newContainer(name, image, c.cli)
}
