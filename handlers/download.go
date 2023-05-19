package handlers

import (
	"fmt"
	"os"
	"rocat/api"
	"rocat/database"
	"rocat/structs"
	"strconv"

	"github.com/urfave/cli/v2"
)

func Download(cCtx *cli.Context, cookie string, downloads database.Storage) error {
	amount, err := strconv.ParseInt(cCtx.String("amount"), 0, 16)

	if err != nil {
		fmt.Println("Please enter a valid clothing limit using the `--amount` flag.")
		return nil
	}

	if amount > 120 {
		fmt.Println("Maximum '--amount' is 120")
		return nil
	}

	if cloths, err := api.GetCatalogue(56, 1, 120, cookie); err != nil {
		fmt.Println("Unable to fetch catalogue")
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf(`Successfuly fetched %v clothing from the catalogue`, len(cloths)))
		if clothes, err := api.GetClothing(structs.GetClothesRequest{
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
					if templateId, err := api.GetTemplateId(cloth.Id); err != nil {
						fmt.Println(err)
					} else {
						path := fmt.Sprintf(`./downloads/%v`, cloth.Id)
						if _, err := os.Stat(path); err != nil {
							if err := api.DownloadTemplate(fmt.Sprintf(`https://www.roblox.com/library/%v`, templateId), path); err != nil {
								fmt.Println(err)
							} else {
								fmt.Println(fmt.Sprintf(`New Template Downloaded, AssetId: %v, TemplateId: %v, Path: %v`, cloth.Id, path, templateId))
								downloads.SaveRecord(structs.Record{
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
}
