package web

import "github.com/urfave/cli/v2"

var WebCommand = &cli.Command{
	Name:  "webserver",
	Usage: "",
	Action: func(c *cli.Context) error {
		listen := c.String("listen")

		return StartGin(listen)
	},
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "listen",
			Aliases: []string{"l"},
			Usage:   "listen address",
			Value:   ":8080",
		},
	},
}
