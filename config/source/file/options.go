package file

import (
	"context"

	"github.com/view-team/view-server-core/config/source"
)

type filePathKey struct{}

// WithPath sets the path to file
func WithPath(p string) source.Option {
	return func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, filePathKey{}, p)
	}
}
