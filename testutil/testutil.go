package testutil

import (
	"testing"

	"github.com/tdakkota/cliflags"
	"github.com/urfave/cli/v2"
)

func Parse(t *testing.T, builder func(cliflags.Namer) []cli.Flag, args ...string) (c *cli.Context) {
	return ParseFlags(t, builder(cliflags.DefaultNamer{}), args...)
}

func ParseFlags(t *testing.T, flags []cli.Flag, args ...string) (c *cli.Context) {
	args = append([]string{"app"}, args...)

	app := cli.NewApp()
	app.Flags = flags
	app.Action = func(ctxt *cli.Context) error {
		c = ctxt
		return nil
	}
	if err := app.Run(args); err != nil {
		t.Fatal(err)
	}

	return c
}
