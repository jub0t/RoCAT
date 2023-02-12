package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

// Get the cloth template/source
func getTemplateId(assetId int) (string, error) {
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
								return strings.Split(strings.Split(strings.Split(string(body), "<url>")[1], "</url>")[0], "?id=")[1], nil
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

// Download template
// Split from: <div id="current-animation-name"></div>
// Then from: <div class="equipped-marker"></div>
func downloadTemplate(link string, path string) error {
	if req, err := http.NewRequest("GET", link, nil); err != nil {
		return err
	} else {
		// Send Request
		if response, err := http.DefaultClient.Do(req); err != nil {
			fmt.Println(err)
			return err
		} else {
			if body, err := ioutil.ReadAll(response.Body); err != nil {
				return err
			} else {
				template := resizeTemplate(strings.Replace(strings.Split(strings.Split(strings.Split(strings.Split(string(body), `<div id="current-animation-name"></div>`)[1], `<div class="equipped-marker"></div>`)[0], "src=")[1], "'/>")[0], "'", ``, 1))

				if req, err := http.NewRequest("GET", template, nil); err != nil {
					return err
				} else {
					// Send Request
					if response, err := http.DefaultClient.Do(req); err != nil {
						fmt.Println(err)
						return err
					} else {
						if body, err := ioutil.ReadAll(response.Body); err != nil {
							return err
						} else {
							if err := os.WriteFile(path, body, os.ModePerm); err != nil {
								fmt.Println(err)
							} else {
								return nil
							}
						}
					}
				}

			}
		}
	}

	return nil
}
