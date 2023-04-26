package main

import (
	"log"
	"os"
	"{{cookiecutter.module_name}}/config"
	"{{cookiecutter.module_name}}/db"
	"{{cookiecutter.module_name}}/logging"
	"{{cookiecutter.module_name}}/web"

	"github.com/gin-gonic/gin"
)

var logger = logging.GetSugar()

func main() {
	app := cli.App{
		Name:  "{{cookiecutter.module_name}}",
		Usage: "",
		Commands: []*cli.Command{
			web.WebCommand,
			db.MigrateCommand,
			config.ConfigCommand,
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "enable debug mode",
				Value:   false,
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "config filename",
				Value:   "config.yaml",
			},
		},
		Before: func(c *cli.Context) error {
			debug := c.Bool("debug")
			err := logging.InitLogger(debug)
			if err != nil {
				return err
			}
			logger.Infof("debug mode = %v", debug)

			// exit before function and generate initial config file
			if funk.ContainsString(os.Args, "genconfig") {
				return nil
			}

			if debug {
				gin.SetMode(gin.DebugMode)
			} else {
				gin.SetMode(gin.ReleaseMode)
			}

			configFile := c.String("config")
			err = config.InitConfig(configFile)
			if err != nil {
				return cli.Exit("failed to load config", 1)
			}

			err = db.InitPostgres(config.GlobalConfig.DatabaseURL, debug)
			if err != nil {
				return cli.Exit("failed to initial PostgreSQL database", 1)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
