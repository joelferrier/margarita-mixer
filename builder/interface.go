package builder

import (
	"github.com/joelferrier/margarita-mixer/config/builder"
	"github.com/joelferrier/margarita-mixer/config/profile"
)

type Builder interface {
	Configure(builder.Config)
	Setup(profile.Config) error
	Build() error
	Extract() error
	Cleanup() error
}
