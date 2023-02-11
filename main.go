package main

import (
	"errors"
	"fmt"
	"os"
)

// Global Variables
const (
	TargetGroup       = 7830839
	AssetAPI          = `https://assetdelivery.roblox.com/v1/assetId/%v`
	CatalogueBatchAPI = "https://catalog.roblox.com/v1/catalog/items/details"
	GetCatalogueAPI   = `https://catalog.roblox.com/v1/search/items?category=Clothing&pxMin=5&limit=%v&salesTypeFilter=1&sortAggregation=%v&sortType=2&subcategory=%v`
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
		if csrf, err := getCSRF(cookie); err != nil {
			fmt.Println(`Unable to get Csrf Token, please re-check your cookie`)
			panic(err)
		} else {
			if shirts, err := getCatalogue(56, 1, 120); err != nil {
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
						} else {
							fmt.Println("Price is less than 5 for:", cloth.Name)
						}
					}

					test := clothes[0]

					if template, err := getTemplateLink(test.Id); err != nil {
						fmt.Println(err)
					} else {
						path := fmt.Sprintf(`./downloads/%v.png`, test.Id)
						if err := downloadTemplate(template, path); err != nil {
							fmt.Println(err)
						} else {
							fmt.Println(fmt.Sprintf(`Template %s Written To %v`, path, test.Id))
						}
					}
				}
			}
		}
	}
}
