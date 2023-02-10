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
			if catalogue, err := getCatalogue(56, 1, 10); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(catalogue)
				if clothes, err := getClothing([]RequestItems{}, cookie, csrf); err != nil {
					fmt.Println(`Unable to get clothes`)
					fmt.Println(err)
				} else {
					fmt.Println(clothes)
				}
			}
		}
	}
}
