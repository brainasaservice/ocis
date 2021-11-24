package command

import (
	"context"
	"os"

	"github.com/owncloud/ocis/glauth/pkg/config"
	ociscfg "github.com/owncloud/ocis/ocis-pkg/config"
	"github.com/owncloud/ocis/ocis-pkg/log"
	oclog "github.com/owncloud/ocis/ocis-pkg/log"
	"github.com/owncloud/ocis/ocis-pkg/shared"
	"github.com/owncloud/ocis/ocis-pkg/version"
	"github.com/thejerf/suture/v4"
	"github.com/urfave/cli/v2"
)

// Execute is the entry point for the ocis-glauth command.
func Execute(cfg *config.Config) error {
	app := &cli.App{
		Name:     "ocis-glauth",
		Version:  version.String,
		Usage:    "Serve GLAuth API for oCIS",
		Compiled: version.Compiled(),
		Authors: []*cli.Author{
			{
				Name:  "ownCloud GmbH",
				Email: "support@owncloud.com",
			},
		},
		Before: func(c *cli.Context) error {
			cfg.Service.Version = version.String
			return nil
		},
		Commands: []*cli.Command{
			Server(cfg),
			Health(cfg),
		},
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:  "help,h",
		Usage: "Show the help",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:  "version,v",
		Usage: "Print the version",
	}

	return app.Run(os.Args)
}

// NewLogger initializes a service-specific logger instance.
func NewLogger(cfg *config.Config) log.Logger {
	return oclog.LoggerFromConfig("glauth", *cfg.Log)
}

// ParseConfig loads glauth configuration from known paths.
func ParseConfig(c *cli.Context, cfg *config.Config) error {
	conf, err := ociscfg.BindSourcesToStructs("glauth", cfg)
	if err != nil {
		return err
	}

	// provide with defaults for shared logging, since we need a valid destination address for BindEnv.
	if cfg.Log == nil && cfg.Commons != nil && cfg.Commons.Log != nil {
		cfg.Log = &shared.Log{
			Level:  cfg.Commons.Log.Level,
			Pretty: cfg.Commons.Log.Pretty,
			Color:  cfg.Commons.Log.Color,
			File:   cfg.Commons.Log.File,
		}
	} else if cfg.Log == nil && cfg.Commons == nil {
		cfg.Log = &shared.Log{}
	}

	// load all env variables relevant to the config in the current context.
	conf.LoadOSEnv(config.GetEnv(cfg), false)
	bindings := config.StructMappings(cfg)
	return ociscfg.BindEnv(conf, bindings)
}

// SutureService allows for the glauth command to be embedded and supervised by a suture supervisor tree.
type SutureService struct {
	cfg *config.Config
}

// NewSutureService creates a new glauth.SutureService
func NewSutureService(cfg *ociscfg.Config) suture.Service {
	cfg.GLAuth.Commons = cfg.Commons
	return SutureService{
		cfg: cfg.GLAuth,
	}
}

func (s SutureService) Serve(ctx context.Context) error {
	s.cfg.Context = ctx
	if err := Execute(s.cfg); err != nil {
		return err
	}

	return nil
}
