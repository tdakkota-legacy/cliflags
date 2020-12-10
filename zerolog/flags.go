package zerolog

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/tdakkota/cliflags"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
)

func Flags(e cliflags.Namer) []cli.Flag {
	return []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     e.Name("logging.format"),
			Required: false,
			Value:    "human",
			Usage:    "logging format(json/human)",
			EnvVars:  e.Env("LOG_FORMAT"),
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     e.Name("logging.level"),
			Required: false,
			Value:    "debug",
			Usage:    "logging level",
			EnvVars:  e.Env("LOG_LEVEL"),
		}),
	}
}

func Build(c *cli.Context) zerolog.Logger {
	level, err := zerolog.ParseLevel(c.String("logging.level"))
	if err != nil {
		level = zerolog.DebugLevel
	}

	var writer io.Writer
	switch c.String("logging.format") {
	case "human":
		writer = zerolog.ConsoleWriter{
			TimeFormat: time.RFC822,
			Out:        os.Stdout,
		}
	case "json":
		writer = os.Stdout
	default:
		writer = os.Stdout
	}

	return zerolog.New(writer).
		Level(level).
		With().
		Caller().
		Timestamp().
		Logger()
}
