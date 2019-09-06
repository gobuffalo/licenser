package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/licenser/genny/licenser"
	"github.com/spf13/cobra"
)

var rootOpts = struct {
	*licenser.Options
	dryRun bool
}{
	Options: &licenser.Options{
		Year: time.Now().Year(),
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "licenser",
	Short: "generates license files",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := genny.WetRunner(context.Background())
		if rootOpts.dryRun {
			r = genny.DryRunner(context.Background())
		}
		opts := rootOpts.Options
		if err := r.WithNew(licenser.New(opts)); err != nil {
			return err
		}
		return r.Run()
	},
}

func main() {
	rootCmd.Flags().BoolVarP(&rootOpts.dryRun, "dry-run", "d", false, "run the generator without creating files or running commands")
	rootCmd.Flags().StringVarP(&rootOpts.Author, "author", "a", "", "author's name")
	rootCmd.Flags().StringVarP(&rootOpts.Name, "license", "l", "mit", fmt.Sprintf("choose a license from: [%s]", strings.Join(licenser.Available, ", ")))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
