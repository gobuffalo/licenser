package licenser

import (
	"fmt"
	"path/filepath"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/plushgen"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
)

var Available []string
var box packd.Box

func init() {
	box = packr.New("github.com/gobuffalo/licenser/genny/licenser/templates", "../licenser/templates")
	box.Walk(func(path string, f packd.File) error {
		name := filepath.Base(path)
		Available = append(Available, name)
		return nil
	})
}

func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()
	if err := opts.Validate(); err != nil {
		return g, err
	}

	body, err := box.Find(opts.Name)
	if err != nil {
		return g, fmt.Errorf("could not find a license named %s", opts.Name)
	}
	g.File(genny.NewFileB("LICENSE.plush", body))

	ctx := plush.NewContext()
	ctx.Set("opts", opts)
	g.Transformer(plushgen.Transformer(ctx))
	return g, nil
}
