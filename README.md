# cliflags

[![Go](https://github.com/tdakkota/cliflags/workflows/Go/badge.svg)](https://github.com/tdakkota/cliflags/actions)
[![Documentation](https://godoc.org/github.com/tdakkota/cliflags?status.svg)](https://pkg.go.dev/github.com/tdakkota/cliflags)
[![license](https://img.shields.io/github/license/tdakkota/cliflags.svg?maxAge=2592000)](https://github.com/tdakkota/cliflags/blob/master/LICENSE)

Ready to use flag sets for [`github.com/urfave/cli`](https://github.com/urfave/cli).

## Example for `vksdk`

```go
package main

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

func main() {
	if err := NewApp().Run(os.Args); err != nil {
		_, _ = os.Stdout.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
```