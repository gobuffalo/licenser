package licenser

import (
	"os/user"
	"strings"
	"time"
)

type Options struct {
	Name   string
	Year   int
	Author string
}

func (opts *Options) Validate() error {
	opts.Name = strings.TrimSpace(opts.Name)
	if opts.Name == "" {
		opts.Name = "mit"
	}
	opts.Name = strings.ToLower(opts.Name)
	if opts.Year == 0 {
		opts.Year = time.Now().Year()
	}

	if opts.Author == "" {
		u, err := user.Current()
		if err != nil {
			return err
		}
		opts.Author = strings.TrimSpace(u.Name)
		if opts.Author == "" {
			opts.Author = strings.TrimSpace(u.Username)
		}
	}

	return nil
}
