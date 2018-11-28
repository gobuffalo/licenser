package licenser

import (
	"path/filepath"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/plushgen"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/pkg/errors"
)

var Available []string
var box packd.Box

func init() {
	box = packr.New("genny:licenser", "../licenser/templates")
	box.Walk(func(path string, f packd.File) error {
		name := filepath.Base(path)
		Available = append(Available, name)
		return nil
	})
}

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()
	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}

	body, err := box.Find(opts.Name)
	if err != nil {
		return g, errors.Errorf("could not find a license named %s", opts.Name)
	}
	g.File(genny.NewFileB("LICENSE.plush", body))

	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	return g, nil
}
