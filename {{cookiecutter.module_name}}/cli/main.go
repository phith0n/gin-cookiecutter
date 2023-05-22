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
				Usage:   "config key or config file name",
			},
		},
		Before: func(c *cli.Context) error {
			configFile := c.String("config")
			err := config.InitConfig(configFile)
			if err != nil {
				return cli.Exit("failed to load config", 1)
			}

			debug := c.Bool("debug")
			if debug {
				config.GlobalConfig.Debug = true
			}

			err = logging.InitLogger(config.GlobalConfig.Debug)
			if err != nil {
				return err
			}
			logger.Infof("debug mode = %v", config.GlobalConfig.Debug)

			// exit before function and generate initial config file
			if funk.ContainsString(os.Args, "genconfig") {
				return nil
			}

			if config.GlobalConfig.Debug {
				gin.SetMode(gin.DebugMode)
			} else {
				gin.SetMode(gin.ReleaseMode)
			}

			{% if cookiecutter.database == "mysql" %}
			err = db.InitMysql(config.GlobalConfig.DatabaseURL, debug)
			if err != nil {
				return cli.Exit("failed to initial MySQL database", 1)
			}
			{% else %}
			err = db.InitPostgres(config.GlobalConfig.DatabaseURL, debug)
			if err != nil {
				return cli.Exit("failed to initial PostgreSQL database", 1)
			}
			{% endif %}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
