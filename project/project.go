package project

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/BurntSushi/toml"
	"github.com/joelferrier/margarita-mixer/config/builder"
	"github.com/joelferrier/margarita-mixer/config/profile"
	"github.com/joelferrier/margarita-mixer/config/repository"
)

const (
	// DefaultBuilderName defines default builder name for a project
	DefaultBuilderName    = "docker"
	DefaultBuilderUrl     = "unix:///var/run/docker.sock"
	DefaultBuilderTimeout = 1800

	// DefaultRepositoryName defines default repository name for a project
	DefaultRepositoryName    = "default"
	DefaultRepositoryGPGSign = false

	// DefaultProfileName defines a default build profile name for a project
	DefaultProfileName           = "amzn1"
	DefaultProfileImage          = "amazonlinux:1"
	DefaultProfilePackageManager = "yum"
)

var (
	DefaultProfileHeaderPackages = []string{"kernel-devel.x86_64"}
	DefaultProfileDependencies   = []string{"git", "gcc", "make"}
)

//Config for project
type Project struct {
	Path         string
	Builders     map[string]builder.Config
	Repositories map[string]repository.Config `toml:"repository"`
	Profiles     map[string]profile.Config    `toml:"profile"`
}

func (p *Project) defaults() {
	p.Builders[DefaultBuilderName] = builder.Config{
		DefaultBuilderUrl,
		DefaultBuilderTimeout,
	}

	p.Repositories[DefaultRepositoryName] = repository.Config{
		DefaultRepositoryGPGSign,
	}

	p.Profiles[DefaultProfileName] = profile.Config{
		DefaultProfileImage,
		DefaultProfilePackageManager,
		DefaultProfileHeaderPackages,
		DefaultProfileDependencies,
	}
}

func (p *Project) Open() error {
	p.defaults()

	configFile := "project.toml"

	_, err := toml.DecodeFile(filepath.Join(p.Path, configFile), &p)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (p *Project) PrintProfiles() {
	var keys []string
	for k := range p.Profiles {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		prof := p.Profiles[k]
		fmt.Printf("[%s]\n%s\n\n", k, prof.ToString())
	}
}

func New() (*Project, error) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	p := Project{
		dir,
		make(map[string]builder.Config),
		make(map[string]repository.Config),
		make(map[string]profile.Config),
	}

	return &p, err
}
