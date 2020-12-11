package td

import (
	"os"
	"path/filepath"

	"github.com/gotd/td/telegram"
	"github.com/tdakkota/cliflags"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
)

func Flags(e cliflags.Namer) []cli.Flag {
	return []cli.Flag{
		// tg
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:     e.Name("tg.app_id"),
			Required: true,
			Usage:    "Telegram app ID",
			Aliases:  []string{"app_id"},
			EnvVars:  e.Env("APP_ID"),
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     e.Name("tg.app_hash"),
			Required: true,
			Usage:    "Telegram app hash",
			Aliases:  []string{"app_hash"},
			EnvVars:  e.Env("APP_HASH"),
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:     e.Name("tg.bot_token"),
			Required: false,
			Usage:    "Telegram bot token",
			Aliases:  []string{"token"},
			EnvVars:  e.Env("BOT_TOKEN"),
		}),
		altsrc.NewPathFlag(&cli.PathFlag{
			Name:    e.Name("tg.session_dir"),
			Usage:   "Telegram session directory",
			Aliases: []string{"session_dir"},
			EnvVars: e.Env("SESSION_DIR"),
		}),
	}
}

func Build(c *cli.Context, logger *zap.Logger, handler telegram.UpdateHandler) (*telegram.Client, error) {
	sessionDir := ""
	if c.IsSet("tg.session_dir") {
		sessionDir = c.String("tg.session_dir")
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			sessionDir = "./.td"
		} else {
			sessionDir = filepath.Join(home, ".td")
		}
	}
	if err := os.MkdirAll(sessionDir, 0600); err != nil {
		return nil, xerrors.Errorf("failed to create session dir: %w", err)
	}

	client, err := telegram.Dial(c.Context, telegram.Options{
		Logger: logger,
		SessionStorage: &telegram.FileSessionStorage{
			Path: filepath.Join(sessionDir, "session.json"),
		},

		// Grab these from https://my.telegram.org/apps.
		// Never share it or hardcode!
		AppID:         c.Int("tg.app_id"),
		AppHash:       c.String("tg.app_hash"),
		UpdateHandler: handler,
	})
	if err != nil {
		return nil, xerrors.Errorf("failed to dial: %w", err)
	}

	auth, err := client.AuthStatus(c.Context)
	if err != nil {
		return nil, xerrors.Errorf("failed to get auth status: %w", err)
	}

	logger.With(zap.Bool("authorized", auth.Authorized)).Info("Auth status")
	if !auth.Authorized {
		if err := client.BotLogin(c.Context, c.String("tg.bot_token")); err != nil {
			return nil, xerrors.Errorf("failed to perform bot login: %w", err)
		}
		logger.Info("Bot login ok")
	}

	return client, nil
}
