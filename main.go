package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"rocat/database"
	"rocat/handlers"
	"rocat/modules"
	_ "rocat/modules"
	"time"

	"github.com/urfave/cli/v2"
)

// Main Function
func main() {
	modules.InitFiles([]string{"./cookie.txt"})
	modules.InitDirs([]string{"./downloads", "./store", "./temp"})

	start_time := time.Now().Unix()
	rand.Seed(start_time)

	downloads, err := database.New("./store/downloads")
	if err != nil {
		fmt.Println(err)
	}

	uploads, err := database.New("./store/uploads")
	if err != nil {
		fmt.Println(err)
	}

	if cookie_file, err := os.ReadFile("cookie.txt"); err != nil {
		fmt.Println(`Unable to get cookie, please make sure you have a 'cookie.txt' file.`)
		panic(err)
	} else {
		cookie := string(cookie_file[:])
		app := &cli.App{
			Name:  "RoCAT",
			Usage: "Roblox clothing automation tool.",
			Commands: []*cli.Command{
				{
					Name:  "info",
					Usage: "Display information about the cli.",
					Action: func(cCtx *cli.Context) error {
						// Database & Downloaded info
						return nil
					},
				},
				{
					Name:  "whoami",
					Usage: "Display bot information using the cookie.",
					Action: func(cCtx *cli.Context) error {
						handlers.WhoAmI(cCtx, cookie)
						return nil
					},
				},
				{
					Name:  "download",
					Usage: "Download classic clothing from roblox catalogue and save them for later upload",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "type",
							Aliases: []string{"t"},
							Usage:   "Clothing type, shirts/pants.",
						},
						&cli.IntFlag{
							Name:    "amount",
							Aliases: []string{"a"},
							Usage:   "Number of clothing templates to download, maximum 120.",
						},
					},
					Action: func(cCtx *cli.Context) error {
						handlers.Download(cCtx, cookie, downloads)
						return nil
					},
				},
				{
					Name:  "upload",
					Usage: "Upload the stored clothing to your account (or group).",
					Flags: []cli.Flag{
						&cli.IntFlag{
							Name:    "groupId",
							Aliases: []string{"gid"},
							Usage:   "Id of the group you want the clothes to upload to.",
						},
						&cli.IntFlag{
							Name:    "limit",
							Aliases: []string{"l"},
							Usage:   "Maximum amount of clothing you want to upload.",
						},
						&cli.BoolFlag{
							Name:    "seo",
							Aliases: []string{"s"},
							Usage:   "Use an algorithm to generate a description that'll help your clothes sell better.",
						},
					},
					Action: func(cCtx *cli.Context) error {
						handlers.Upload(cCtx, cookie, downloads, uploads)
						return nil
					},
				},
			},
		}

		if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
		}
	}
}
