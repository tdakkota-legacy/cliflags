package vksdk_test

import (
	"fmt"
	"os"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/tdakkota/cliflags"
	"github.com/tdakkota/cliflags/vksdk"
	"github.com/urfave/cli/v2"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (app *App) cli() *cli.App {
	cliApp := &cli.App{
		Name:  "example app",
		Flags: vksdk.Flags(cliflags.DefaultNamer{}),
		Action: func(c *cli.Context) error {
			b, err := vksdk.Build(c)
			if err != nil {
				return err
			}

			p := params.NewUsersGetBuilder()
			p.WithContext(c.Context)
			r, err := b.Complete().UsersGet(p.Params)
			if err != nil {
				return err
			}

			fmt.Println(r)
			return nil
		},
	}

	return cliApp
}

func (app *App) Run(args []string) error {
	return app.cli().Run(args)
}

func Example_vksdk() {
	if err := NewApp().Run(os.Args); err != nil {
		_, _ = os.Stdout.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
