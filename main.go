package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Global Variables
const (
	TargetGroup       = 7830839
	AssetAPI          = `https://assetdelivery.roblox.com/v1/assetId/%v`
	CatalogueBatchAPI = "https://catalog.roblox.com/v1/catalog/items/details"
	GetCatalogueAPI   = `https://catalog.roblox.com/v1/search/items?category=Clothing&limit=%v&salesTypeFilter=1&sortAggregation=%v&sortType=2&subcategory=%v&minPrice=5`
)

// Main Function
func main() {
	if _, err := os.Stat("./downloads"); err == nil {
	} else if errors.Is(err, os.ErrNotExist) {
		os.MkdirAll("./downloads", os.ModePerm)
	} else {
		os.MkdirAll("./downloads", os.ModePerm)
	}

	if cookie_file, err := os.ReadFile("cookie.txt"); err != nil {
		fmt.Println(`Unable to get cookie, please make sure you have a 'cookie.txt' file.`)
		panic(err)
	} else {
		cookie := string(cookie_file[:])

		app := &cli.App{
			Name:  "RoCat",
			Usage: "Roblox clothing automation tool.",
			Commands: []*cli.Command{
				{
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "type",
							Aliases: []string{"t"},
							Usage:   "Clothing type, shirts/pants.",
						},
					},
					Name:    "download",
					Aliases: []string{"dw"},
					Usage:   "Download classic clothing from roblox catalogue and save them for later upload",
					Action: func(cCtx *cli.Context) error {
						fmt.Println(cCtx.App.Flags)
						fmt.Println(cCtx.Args())

						if csrf, err := getCSRF(cookie); err != nil {
							fmt.Println(`Unable to get Csrf Token, please re-check your cookie`)
							panic(err)
						} else {
							if shirts, err := getCatalogue(56, 1, 10); err != nil {
								fmt.Println(err)
							} else {
								if clothes, err := getClothing(GetClothesRequest{
									Items: shirts,
								}, cookie, csrf); err != nil {
									fmt.Println(err)
								} else {
									for i := 0; i < len(clothes); i++ {
										cloth := clothes[i]

										if cloth.Price >= 5 {
											fmt.Println(cloth.Id, cloth.ProductId, cloth.Price, cloth.Name)
											if template, err := getTemplateLink(cloth.Id); err != nil {
												fmt.Println(err)
											} else {
												path := fmt.Sprintf(`./downloads/%v.png`, cloth.Id)
												if _, err := os.Stat(path); err != nil {
													if err := downloadTemplate(template, path); err != nil {
														fmt.Println(err)
													} else {
														fmt.Println(fmt.Sprintf(`Template %s Written To %v`, path, cloth.Id))
													}
												}
											}
										}
									}
								}
							}
						}

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
