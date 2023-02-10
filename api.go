package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Get Clothing Information
func getClothing(items GetClothesRequest, cookie string, csrf string) ([]ResponseItems, error) {
	if body, err := json.Marshal(items); err != nil {
		return nil, err
	} else {
		if req, err := http.NewRequest("POST", CatalogueBatchAPI, bytes.NewReader(body)); err != nil {
			fmt.Println("Request Agent Error")
			return nil, err
		} else {
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set(".ROBLOSECURITY", cookie)
			req.Header.Set("x-csrf-token", csrf)

			if response, err := http.DefaultClient.Do(req); err != nil {
				fmt.Println("Response Error")
				return nil, err
			} else {
				if body, err := ioutil.ReadAll(response.Body); err != nil {
					return nil, err
				} else {
					var catalogue GetClothesResponse
					if err := json.Unmarshal(body, &catalogue); err != nil {
						return nil, err
					} else {
						return catalogue.Data, nil
					}
				}
			}
		}
	}
}

func getCSRF(cookie string) (string, error) {
	if req, err := http.NewRequest("POST", "https://auth.roblox.com/v2/login", nil); err != nil {
		return "", err
	} else {
		req.Header.Set(".ROBLOSECURITY", cookie)

		// Send Request
		if response, err := http.DefaultClient.Do(req); err != nil {
			fmt.Println(err)
			return "", err
		} else {
			// Get CSRF token from headers
			token := response.Header.Get("x-csrf-token")
			return token, nil
		}
	}
}

// Get Shirt/Pants from catalogue
func getCatalogue(sub int, agg int, limit int) ([]CatalogueItem, error) {
	url := fmt.Sprintf(GetCatalogueAPI, limit, agg, sub)
	if req, err := http.NewRequest("GET", url, nil); err != nil {
		return []CatalogueItem{}, err
	} else {
		req.Header.Set("Content-Type", "application/json")

		// Send Request
		if response, err := http.DefaultClient.Do(req); err != nil {
			fmt.Println(err)
			return []CatalogueItem{}, err
		} else {
			if body, err := ioutil.ReadAll(response.Body); err != nil {
				return []CatalogueItem{}, err
			} else {
				var catalogue CatalogueResponse

				if err := json.Unmarshal(body, &catalogue); err != nil {
					return []CatalogueItem{}, err
				} else {
					return catalogue.Data, nil
				}
			}
		}
	}
}

// Fetch asset information
func getAssetInfo(assetId int) {

}

// Get the cloth template/source
func getTemplateLink(assetId int) (string, error) {
	if request, err := http.NewRequest("GET", fmt.Sprintf(AssetAPI, assetId), nil); err != nil {
		fmt.Println(err)
	} else {
		request.Header.Set("Content-Type", "application/json")

		// Send Request
		if response, err := http.DefaultClient.Do(request); err != nil {
			fmt.Println(err)
			return "", err
		} else {
			if body, err := ioutil.ReadAll(response.Body); err != nil {
				return "", err
			} else {
				var asset AssetInfo

				if err := json.Unmarshal(body, &asset); err != nil {
					return "", err
				} else {
					location := asset.Location

					if request, err := http.NewRequest("GET", location, nil); err != nil {
						fmt.Println(err)
					} else {
						request.Header.Set("Content-Type", "application/json")

						// Send Request
						if response, err := http.DefaultClient.Do(request); err != nil {
							fmt.Println(err)
							return "", err
						} else {
							if body, err := ioutil.ReadAll(response.Body); err != nil {
								return "", err
							} else {
								return fmt.Sprintf(`https://www.roblox.com/library/%v`, strings.Split(strings.Split(strings.Split(string(body), "<url>")[1], "</url>")[0], "?id=")[1]), nil
							}
						}
					}
					return "", nil
				}
			}
		}
	}

	return "", nil
}
