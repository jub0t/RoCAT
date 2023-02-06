package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	CatalogueBatchAPI = "https://catalog.roblox.com/v1/catalog/items/details"
	GetCatalogueAPI   = `https://catalog.roblox.com/v1/search/items?category=Clothing&limit=%v&salesTypeFilter=1&sortAggregation=%v&sortType=2&subcategory=%v`
)

// Main Function
func main() {
	if clothes, err := getClothing([]RequestItems{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(clothes)
	}
}

// Get Shirt/Pants from catalogue
func getCatalogue(sub int, agg int, limit int) {
	url := fmt.Sprintf(GetCatalogueAPI, limit, agg, sub)
	fmt.Println(url)
}

func getClothing(items []RequestItems) ([]ResponseItems, error) {
	// Convert items to JSON
	// Send request to CatalogueBatchAPI
	// Convert response to CatalogueResponseItem

	if body, err := json.Marshal(items); err != nil {
		return nil, err
	} else {
		reader := ioutil.NopCloser(bytes.NewReader(body))
		if req, err := http.NewRequest("POST", CatalogueBatchAPI, reader); err != nil {
			fmt.Println("Request Agent Error`")
			return nil, err
		} else {
			req.Header.Set("Content-Type", "application/json")

			if response, err := http.DefaultClient.Do(req); err != nil {
				fmt.Println("Response Error")
				return nil, err
			} else {
				if body, err := ioutil.ReadAll(response.Body); err != nil {
					return nil, err
				} else {
					fmt.Println(fmt.Sprintf("%+v", string(body)))
					var catalogue []ResponseItems
					if err := json.Unmarshal(body, &catalogue); err != nil {
						return nil, err
					} else {
						return catalogue, nil
					}
				}
			}
		}
	}
}

func getCSRF(cookie string) (*string, error) {
	// Get CSRF token from cookie
	// Return CSRF token

	if req, err := http.NewRequest("GET", "https://www.roblox.com", nil); err != nil {
		return nil, err
	} else {
		req.Header.Set(".ROBLOSECURITY", cookie)

		if response, err := http.DefaultClient.Do(req); err != nil {
			return nil, err
		} else {
			// Get CSRF token from headers
			if csrf, ok := response.Header["X-CSRF-TOKEN"]; ok {
				return &csrf[0], nil
			} else {
				return nil, fmt.Errorf("No CSRF token found in response headers")
			}
		}
	}
}
