package main

import (
	"log"
	"os"

	"github.com/anyshare/anyshare-cli/cmd"
	"github.com/anyshare/anyshare-common/mongodb"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

func main() {
	godotenv.Load()
	mongodb.Connect()
	defer mongodb.Disconnect()
	app := &cli.App{
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name: "user",
				Subcommands: []*cli.Command{
					{
						Name:   "default",
						Usage:  "create default admin user",
						Action: cmd.CreateDefaultUser,
					},
					{
						Name:   "fake",
						Usage:  "gen 100 fake user",
						Action: cmd.FakeUser,
					},
				},
			},
			{
				Name: "mongodb",
				Subcommands: []*cli.Command{
					{
						Name:   "index",
						Usage:  "create mongodb indexes",
						Action: cmd.CreateMongodbIndexes,
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
