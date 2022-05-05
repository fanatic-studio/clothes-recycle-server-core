package memory

import (
	"github.com/fanatic-studio/clothes-recycle-server-core/config/loader"
	"github.com/fanatic-studio/clothes-recycle-server-core/config/reader"
	"github.com/fanatic-studio/clothes-recycle-server-core/config/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
