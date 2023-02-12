package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

// Global Variables
const (
	AssetAPI          = `https://assetdelivery.roblox.com/v1/assetId/%v`
	CatalogueBatchAPI = "https://catalog.roblox.com/v1/catalog/items/details"
	UploadAPI         = `https://itemconfiguration.roblox.com/v1/avatar-assets/Shirt/upload`
	GetCatalogueAPI   = `https://catalog.roblox.com/v1/search/items?category=Clothing&limit=%v&salesTypeFilter=1&sortAggregation=%v&sortType=2&subcategory=%v&minPrice=5`
)

// Types
const (
	TypeShirt = 56
	TypePant  = 57
)

// Main Function
func main() {
	initDirs([]string{"./downloads", "./store"})
	initFiles([]string{"./store/database"})

	storage, err := New("./store/database")
	if err != nil {
		fmt.Println(err)
	}

	storage.SaveRecord(Record{
		Type: TypeShirt,
		Name: "Test",
		Id:   1,
	})

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
					Name:    "download",
					Aliases: []string{"dw"},
					Usage:   "Download classic clothing from roblox catalogue and save them for later upload",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "type",
							Aliases: []string{"t"},
							Usage:   "Clothing type, shirts/pants.",
						},
						&cli.IntFlag{
							Name:    "amount",
							Aliases: []string{"a"},
							Usage:   "Number of clothing templates to download, allowed: 10, 28 & 120",
						},
					},
					Action: func(cCtx *cli.Context) error {
						amount, err := strconv.ParseInt(cCtx.String("amount"), 0, 16)

						if err != nil {
							fmt.Println("Please enter a valid clothing limit using the `--limit` flag")
							return nil
						}

						if csrf, err := getCSRF(cookie); err != nil {
							fmt.Println(`Unable to get Csrf Token, please re-check your cookie`)
							panic(err)
						} else {
							if shirts, err := getCatalogue(56, 1, int(amount)); err != nil {
								fmt.Println(err)
							} else {
								if clothes, err := getClothing(GetClothesRequest{
									Items: shirts,
								}, cookie, csrf); err != nil {
									fmt.Println(err)
								} else {
									for i := 0; i < len(clothes); i++ {
										cloth := clothes[i]

										// Avoid re-uploading free clothes
										if cloth.Price >= 5 {
											if template, err := getTemplateLink(cloth.Id); err != nil {
												fmt.Println(err)
											} else {
												path := fmt.Sprintf(`./downloads/%v.png`, cloth.Id)
												if _, err := os.Stat(path); err != nil {
													if err := downloadTemplate(template, path); err != nil {
														fmt.Println(err)
													} else {
														fmt.Println(fmt.Sprintf(`Template with Id %s Written To %v`, path, cloth.Id))
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
				{
					Name:    "start",
					Aliases: []string{"st"},
					Usage:   "Start uploading the stored clothing.",
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
					},
					Action: func(cCtx *cli.Context) error {
						limit, err := strconv.ParseInt(cCtx.String("limit"), 0, 16)

						if err != nil {
							fmt.Println("Please enter a valid clothing limit using the `--limit` flag")
							return nil
						}

						group_id, err := strconv.ParseInt(cCtx.String("groupId"), 0, 32)

						if err != nil {
							fmt.Println("Please enter a valid group Id with `--groupId` flag")
							return nil
						}

						fmt.Println(limit)
						fmt.Println(group_id)

						entries, err := os.ReadDir("./downloads")

						if err != nil {
							fmt.Println(err)
						}

						fmt.Println(fmt.Sprintf(`Loaded %v Clothing Templates from Storage`, len(entries)))

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
