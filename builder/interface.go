package builder

import "github.com/joelferrier/margarita-mixer/config/builder"

type Builder interface {
	Configure(builder.Config)
	Setup() error
	Build() error
	Extract() error
	Cleanup() error
}
