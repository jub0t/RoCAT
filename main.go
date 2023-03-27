package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/urfave/cli/v2"
)

// Global Variables
const (
	AssetAPI          = `https://assetdelivery.roblox.com/v1/assetId/%v`
	CatalogueBatchAPI = "https://catalog.roblox.com/v1/catalog/items/details"
	Alpha             = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	UploadAPI         = `https://itemconfiguration.roblox.com/v1/avatar-assets/Shirt/upload`
	GetCatalogueAPI   = `https://catalog.roblox.com/v1/search/items?category=Clothing&limit=%v&salesTypeFilter=1&sortAggregation=%v&sortType=2&subcategory=%v&minPrice=5`
)

// generates a random string of fixed size
func srand(size int) string {
	buf := make([]byte, size)

	for i := 0; i < size; i++ {
		buf[i] = Alpha[rand.Intn(len(Alpha))]
	}

	return string(buf)
}

// Main Function
func main() {
	initFiles([]string{"./cookie.txt"})
	initDirs([]string{"./downloads", "./store", "./temp"})

	start_time := time.Now().Unix()
	rand.Seed(start_time)

	downloads, err := New("./store/downloads")
	if err != nil {
		fmt.Println(err)
	}

	uploads, err := New("./store/uploads")
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
					Usage: "Uses your cookie from the file and fetches account/bot info.",
					Action: func(cCtx *cli.Context) error {
						user, err := getUserInfo(cookie)

						if err != nil {
							fmt.Println(`Unable to fetch user info, please re-check your cookie.`)
						}

						if user.UserId > 0 {
							fmt.Println(fmt.Sprintf("Username: %v\nUserId: %v\nRobux: %v\nHas Premium: %v\nHas Builders Club: %v", user.UserName, user.UserId, user.RobuxBalance, user.IsPremium, user.IsAnyBuildersClubMember))
						}

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
						amount, err := strconv.ParseInt(cCtx.String("amount"), 0, 16)

						if err != nil {
							fmt.Println("Please enter a valid clothing limit using the `--amount` flag.")
							return nil
						}

						if amount > 120 {
							fmt.Println("Maximum '--amount' is 120")
							return nil
						}

						if cloths, err := getCatalogue(56, 1, 120, cookie); err != nil {
							fmt.Println("Unable to fetch catalogue")
							fmt.Println(err)
						} else {
							fmt.Println(fmt.Sprintf(`Successfuly fetched %v clothing from the catalogue`, len(cloths)))
							if clothes, err := getClothing(GetClothesRequest{
								Items: cloths,
							}, cookie); err != nil {
								fmt.Println("Unable to clothing info")
								fmt.Println(err)
							} else {
								fmt.Println(fmt.Sprintf(`Successfuly fetched asset information for %v clothes`, amount))

								for i := 0; i < len(clothes); i++ {
									cloth := clothes[i]

									if i >= int(amount) {
										fmt.Println(fmt.Sprintf("Successfuly downloaded %v/%v clothing templates", amount, i))
										break
									}

									// Avoid re-uploading free clothes
									if cloth.Price >= 5 {
										if templateId, err := getTemplateId(cloth.Id); err != nil {
											fmt.Println(err)
										} else {
											path := fmt.Sprintf(`./downloads/%v`, cloth.Id)
											if _, err := os.Stat(path); err != nil {
												if err := downloadTemplate(fmt.Sprintf(`https://www.roblox.com/library/%v`, templateId), path); err != nil {
													fmt.Println(err)
												} else {
													fmt.Println(fmt.Sprintf(`New Template Downloaded, AssetId: %v, TemplateId: %v, Path: %v`, cloth.Id, path, templateId))
													downloads.SaveRecord(Record{
														Type: cloth.ItemType,
														Name: cloth.Name,
														Id:   cloth.Id,
													})
												}
											}
										}
									} else {
										fmt.Println(fmt.Sprintf(`Skipping %v, Reason: Price less than 5`, cloth.Id))
									}
								}
							}
						}

						return nil
					},
				},
				{
					Name:  "start",
					Usage: "Start uploading the stored clothing.",
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
						user, err := getUserInfo(cookie)

						if err != nil {
							fmt.Println(`Unable to fetch user info, please re-check your cookie.`)
						}

						fmt.Println(fmt.Sprintf("Logged in as %v(%v), Account Balance: %v", user.UserName, user.UserId, user.RobuxBalance))

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

						seo := cCtx.Bool("seo")
						entries, err := os.ReadDir("./downloads")

						if err != nil {
							fmt.Println(err)
						}

						fmt.Println(fmt.Sprintf(`Loaded %v Clothing Templates from Storage`, len(entries)))

						for i := 0; i < len(entries); i++ {
							if i > int(limit) {
								fmt.Println(fmt.Sprintf(`Limit(%v) reached.`, limit))
								break
							}

							file := entries[i]
							file_name, err := strconv.ParseInt(file.Name(), 0, 64)

							if err != nil {
								fmt.Println(`Cannot parse id, skipping`)
								continue
							}

							// Template has already been uploaded and recorded
							if uploads.RecordExists(int(file_name)) {
								fmt.Println(fmt.Sprintf(`Template(%v) Has Alread Been Uploaded`, file_name))
								continue
							} else {
								// Get template's information
								info := downloads.GetRecord(int(file_name))

								// If it's valid
								if info.Id > 0 {
									fmt.Println(fmt.Sprintf(`Uploading %v`, info.Name))
									if err := uploadTemplate(cookie, info.Name, int(group_id), "Group", info.Id, 5, seo); err != nil {
										fmt.Println(err)
									} else {
										fmt.Println(fmt.Sprintf(`Template %v (%v) Successfuly Uploaded`, info.Name, info.Id))
									}

									// save upload record
									// uploads.SaveRecord(info)
								}
							}
						}

						fmt.Println(fmt.Sprintf(`Successful uploaded %v clothes.`, min(int(limit), len(entries))))

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
