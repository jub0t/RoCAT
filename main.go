package main

import (
	"fmt"
	"os"
)

// Global Variables
const (
	CatalogueBatchAPI = "https://catalog.roblox.com/v1/catalog/items/details"
	GetCatalogueAPI   = `https://catalog.roblox.com/v1/search/items?category=Clothing&limit=%v&salesTypeFilter=1&sortAggregation=%v&sortType=2&subcategory=%v`
)

// Main Function
func main() {
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
							fmt.Println(cloth.Id, cloth.Name, cloth.Price, cloth.ProductId)
						}
					}
				}
			}
		}
	}
}
