package handlers

import (
	"fmt"
	"os"
	"rocat/api"
	"rocat/database"
	"rocat/modules"
	"strconv"

	"github.com/urfave/cli/v2"
)

func Upload(cCtx *cli.Context, cookie string, downloads database.Storage, uploads database.Storage) error {
	user, err := api.GetUserInfo(cookie)

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
				if err := api.UploadTemplate(cookie, info.Name, int(group_id), "Group", info.Id, 5, seo); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(fmt.Sprintf(`Template %v (%v) Successfuly Uploaded`, info.Name, info.Id))
				}

				// save upload record
				// uploads.SaveRecord(info)
			}
		}
	}

	fmt.Println(fmt.Sprintf(`Successful uploaded %v clothes.`, modules.Min(int(limit), len(entries))))

	return nil
}
