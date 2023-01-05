package db

import "github.com/urfave/cli/v2"

var MigrateCommand = &cli.Command{
	Name:  "migrate",
	Usage: "migrate database changes automatically",
	Action: func(c *cli.Context) error {
		return DB.AutoMigrate(
			// &UserModel{},
		)
	},
	Flags: []cli.Flag{},
}
