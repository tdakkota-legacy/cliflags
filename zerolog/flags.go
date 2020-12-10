package vksdk

import (
	"errors"

	"github.com/SevereCloud/vksdk/v2"

	"github.com/tdakkota/cliflags"

	"github.com/tdakkota/vksdkutil/v2"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func Flags(e cliflags.Namer) []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     e.Name("vk.server"),
			Required: false,
			Value:    "https://api.vk.com/method/",
			Usage:    "VK API Method URL",
			EnvVars:  e.Env("VK_SERVER"),
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     e.Name("vk.user_agent"),
			Required: false,
			Value:    "vksdk/" + vksdk.Version + " (+https://github.com/SevereCloud/vksdk)",
			Usage:    "VK API Client useragent",
			EnvVars:  e.Env("VK_USER_AGENT"),
		}),
		altsrc.NewStringSliceFlag(&cli.StringSliceFlag{
			Name:     e.Name("vk.tokens"),
			Aliases:  []string{"tokens"},
			Required: true,
			Usage:    "VK API tokens",
			EnvVars:  e.Env("VK_TOKENS"),
		}),
	}
}

var ErrAtLeastOneTokenExpected = errors.New("at least one token expected")

func Build(c *cli.Context) (sdkutil.SDKBuilder, error) {
	tokens := c.StringSlice("vk.tokens")
	if len(tokens) < 1 {
		return sdkutil.SDKBuilder{}, ErrAtLeastOneTokenExpected
	}

	sdk := sdkutil.BuildSDK(tokens[0], tokens[1:]...).
		WithUserAgent(c.String("vk.user_agent")).
		WithMethodURL(c.String("vk.server"))
	return sdk, nil
}
