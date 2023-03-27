package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
)

// Get Clothing Information
func getClothing(items GetClothesRequest, cookie string) ([]ResponseItems, error) {
	if body, err := json.Marshal(items); err != nil {
		return nil, err
	} else {
		if req, err := http.NewRequest("POST", CatalogueBatchAPI, bytes.NewReader(body)); err != nil {
			return nil, err
		} else {
			if csrf, err := getCSRF(cookie); err != nil {
				fmt.Println(`x-csrf-token fetching failed upon getting clothing`)
				return nil, err
			} else {
				cookie = fmt.Sprintf(".ROBLOSECURITY=%s", cookie)

				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("x-csrf-token", csrf)
				req.Header.Set("cookie", cookie)

				if response, err := http.DefaultClient.Do(req); err != nil {
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
func getCatalogue(sub int, agg int, limit int, cookie string) ([]CatalogueItem, error) {
	if req, err := http.NewRequest("GET", fmt.Sprintf(GetCatalogueAPI, limit, agg, sub), nil); err != nil {
		return []CatalogueItem{}, err
	} else {
		if csrf, err := getCSRF(cookie); err != nil {
			fmt.Println(`x-csrf-token fetching failed upon getting catalogue`)
			panic(err)
		} else {
			req.Header.Set("cookie", fmt.Sprintf(`.ROBLOSECURITY=%v`, cookie))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("x-csrf-token", csrf)

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

// Get User's Balance
func getBalance(cookie string, csrf string, user_id int) (int, error) {
	if req, err := http.NewRequest("GET", fmt.Sprintf(`https://economy.roblox.com/v1/users/%v/currency`, user_id), nil); err != nil {
		fmt.Println("Request Agent Error")
		return 0, err
	} else {
		req.Header.Set("cookie", fmt.Sprintf(`.ROBLOSECURITY=%v`, cookie))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-csrf-token", csrf)

		if response, err := http.DefaultClient.Do(req); err != nil {
			fmt.Println("Response Error")
			return 0, err
		} else {

			if body, err := ioutil.ReadAll(response.Body); err != nil {
				return 0, err
			} else {
				var resp AccountBalanceResponse

				if err := json.Unmarshal(body, &resp); err != nil {
					return 0, err
				} else {
					return resp.Robux, nil
				}
			}
		}
	}
}

// Get the user's information by cookie
func getUserInfo(cookie string) (UserInfo, error) {
	if req, err := http.NewRequest("GET", `https://www.roblox.com/mobileapi/userinfo`, nil); err != nil {
		fmt.Println("Request Agent Error")
		return UserInfo{}, err
	} else {
		if csrf, err := getCSRF(cookie); err != nil {
			fmt.Println(`x-csrf-token fetching failed upon getting user info`)
			panic(err)
		} else {
			req.Header.Set("cookie", fmt.Sprintf(`.ROBLOSECURITY=%v`, cookie))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("x-csrf-token", csrf)

			if response, err := http.DefaultClient.Do(req); err != nil {
				return UserInfo{}, err
			} else {
				if body, err := ioutil.ReadAll(response.Body); err != nil {
					return UserInfo{}, err
				} else {
					var resp UserInfo

					if err := json.Unmarshal(body, &resp); err != nil {
						return UserInfo{}, err
					} else {
						return resp, nil
					}
				}
			}
		}
	}
}

// Upload the template to roblox
func uploadTemplate(cookie string, name string, creator_id int, creatorType string, location int, price int, use_seo bool) error {
	file, _ := os.Open(fmt.Sprintf(`./downloads/%v`, location))
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.SetBoundary(randomBoundary())

	image_body := textproto.MIMEHeader{}
	image_body.Set("Content-Type", "image/png")
	image_body.Set("Content-Disposition", fmt.Sprintf(`form-data; name="content"; filename="shirt.png"`))

	part, _ := writer.CreatePart(image_body)

	json_body := textproto.MIMEHeader{}
	json_body.Set("Content-Type", "application/json")
	json_body.Set("Content-Disposition", `form-data; name="config"; filename="blob"`)

	config, _ := writer.CreatePart(json_body)

	captcha_json := textproto.MIMEHeader{}
	json_body.Set("Content-Type", "application/json")
	json_body.Set("Content-Disposition", `form-data; name="captcha-token-request"; filename="captcha-token-request.json"`)

	captcha, _ := writer.CreatePart(captcha_json)

	u_bytes, err := json.Marshal(UploadConfig{
		Name:            name,
		CreatorTargetId: creator_id,
		CreatorType:     creatorType,
		Description:     name,
	})

	c_bytes, err := json.Marshal(CaptchaConfig{
		CaptchToken:     "",
		CaptchaProvider: "",
	})

	if err != nil {
		fmt.Println(err)
	}

	io.Copy(part, file)
	config.Write(u_bytes)
	captcha.Write(c_bytes)
	writer.Close()

	fmt.Println(body)

	// client := api2captcha.NewClient("YOUR_API_KEY")

	// cap := api2captcha.GeeTest{
	// 	GT:        "f2ae6cadcf7886856696502e1d55e00c",
	// 	Challenge: "12345678abc90123d45678ef90123a456b",
	// 	Url:       "https://mysite.com/captcha.html",
	// 	ApiServer: "api-na.geetest.com",
	// }

	// req := cap.ToRequest()
	// code, err := client.Solve(req)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(code)
	// }

	if req, err := http.NewRequest("POST", UploadAPI, body); err != nil {
		return err
	} else {
		if csrf, err := getCSRF(cookie); err != nil {
			fmt.Println(`x-csrf-token fetching failed for upload`)
			panic(err)
		} else {
			req.Header.Set("content-type", fmt.Sprintf(writer.FormDataContentType()))
			req.Header.Set("cookie", fmt.Sprintf(".ROBLOSECURITY=%v", cookie))
			req.Header.Set("referer", `https://create.roblox.com`)
			req.Header.Set("x-csrf-token", csrf)

			if response, err := http.DefaultClient.Do(req); err != nil {
				return err
			} else {
				if resp, err := ioutil.ReadAll(response.Body); err != nil {
					return err
				} else {
					fmt.Println(string(resp))

					return nil
				}
			}
		}

	}
}
