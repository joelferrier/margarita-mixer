package profile

import (
	"fmt"
)

type Config struct {
	Image          string
	PackageManager string   `toml:"package_manager"`
	HeaderPackages []string `toml:"header_packages"`
	Dependencies   []string
}

func (p *Config) ToString() string {
	return fmt.Sprintf(
		`  image: %s
  package manager: %s
  header packages: %s
  dependencies: %s`,
		p.Image,
		p.PackageManager,
		p.HeaderPackages,
		p.Dependencies,
	)
}
