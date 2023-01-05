package config

import (
	"gopkg.in/yaml.v2"
	"os"
	"{{cookiecutter.module_name}}/utils"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

var ConfigCommand = &cli.Command{
	Name:  "genconfig",
	Usage: "generate a initial config file",
	Action: func(c *cli.Context) error {
		var filename = c.Path("filename")
		if utils.FileExists(filename) {
			return cli.Exit("config file already exists", 1)
		}

		config := &Config{
			WebAddr: ":8080",
			DatabaseURL: "postgres://postgres:postgres@127.0.0.1:5432/example",
		}

		data, err := yaml.Marshal(config)
		if err != nil {
			return cli.Exit("failed to generate config file", 1)
		}

		err = os.WriteFile(filename, data, 0o644)
		if err != nil {
			return cli.Exit("failed to write config file", 1)
		}

		return nil
	},
	Flags: []cli.Flag{
		&cli.PathFlag{
			Name:    "filename",
			Usage:   "config file path",
			Aliases: []string{"c"},
			Value:   "config.yaml",
		},
	},
}
