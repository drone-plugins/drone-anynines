package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "anynines plugin"
	app.Usage = "anynines plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "username",
			Usage:  "username for anynines auth",
			EnvVar: "PLUGIN_USERNAME,ANYNINES_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "password for anynines auth",
			EnvVar: "PLUGIN_PASSWORD,ANYNINES_PASSWORD",
		},
		cli.StringFlag{
			Name:   "organization",
			Usage:  "organization on anynines",
			EnvVar: "PLUGIN_ORGANIZATION,ANYNINES_ORGANIZATION",
			Value:  "application/json",
		},
		cli.StringFlag{
			Name:   "space",
			Usage:  "space within anynines organization",
			EnvVar: "PLUGIN_SPACE,ANYNINES_SPACE",
			Value:  "production",
		},
		cli.BoolFlag{
			Name:   "skip-cleanup",
			Usage:  "skip cleanup of workspace",
			EnvVar: "PLUGIN_SKIP_CLEANUP",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Config: Config{
			Username:     c.String("username"),
			Password:     c.String("password"),
			Organization: c.String("organization"),
			Space:        c.String("space"),
			SkipCleanup:  c.Bool("skip-cleanup"),
		},
	}

	if plugin.Config.Username == "" {
		return errors.New("Missing username")
	}

	if plugin.Config.Password == "" {
		return errors.New("Missing password")
	}

	if plugin.Config.Organization == "" {
		return errors.New("Missing organization")
	}

	return plugin.Exec()
}
