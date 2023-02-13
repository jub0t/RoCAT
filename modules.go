package main

import (
	"errors"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	SeoWords = []string{"Fashion", "Style", "Cute", "Beauty", "Pretty",
		"Beautiful", "Geek", "Adorable", "Amazing", "Nice", "Chill", "Gorgeous",
		"Girls", "Girly", "Tomboy", "Design", "Model", "Headphones", "Beats", "Shirt",
		"Tube Top", "Cami", "Camisole", "Bandeau", "Crop Top", "Flannel", "Plaid", "Sequin",
		"Jacket", "Pullover", "Cardigan", "Sweatshirt", "Dress", "Denim", "Shortsa", "Leggings",
		"Jeans", "Skirt", "Pants", "Overalls", "Swimsuit", "Bikini", "Boots", "Combat", "Shoes",
		"Heels", "Converse", "Uggs", "Vans", "Outfit", "Glam", "Red", "Orange", "Yellow", "Green",
		"Blue", "Purple", "Brown", "Pink", "Navy", "White", "Rainbow", "Black", "Floral", "Galaxy",
		"Mustache", "Print", "Fall", "Winter", "Spring", "Summer", "Beach", "Easter", "Christmas",
		"Thanksgiving", "Halloween", "Princess", "Prince", "Queen", "King", "aerotags: adidas",
		"nike", "roblox", "anime", "games shirt", "pants", "top", "red", "blue", "yellow", "cool",
		"rebook", "villain", "pro", "noob", "epic", "long", "hoodie", "purple", "hood", "swear",
		"gratis", "shoes", "camiseta", "pink", "ice", "police", "original", "baby", "limited", "sweet",
		"bills", "kawaii", "lion", "emo", "goth", "y2k"}
)

// Create directories if not exist
func initDirs(paths []string) {
	for i := 0; i < len(paths); i++ {
		path := paths[i]

		if _, err := os.Stat(path); err == nil {
		} else if errors.Is(err, os.ErrNotExist) {
			os.MkdirAll(path, os.ModePerm)
		} else {
			os.MkdirAll(path, os.ModePerm)
		}
	}
}

// Create files if not exist
func initFiles(files []string) {
	for i := 0; i < len(files); i++ {
		file := files[i]

		if _, err := os.Stat(file); err != nil {
			_, e := os.Create(file)

			if e != nil {
				fmt.Println(fmt.Sprintf(`Failed creating file: %v`, file))
			}
		}
	}

}

// The clothing size is 420, we resize it to 512 without losing quality
func resizeTemplate(link string) string {
	return fmt.Sprintf("https://tr.rbxcdn.com/%v/512/512/Image/Png", strings.Split(strings.Split(link, "https://tr.rbxcdn.com/")[1], "/")[0])
}

// Return only the names of fs dirs
func entriesToNames(entries []fs.DirEntry) []string {
	var x []string

	for i := 0; i < len(entries); i++ {
		x = append(x, entries[i].Name())
	}

	return x
}

func contains(s []string, e any) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Map a slice and only return clothes that have not been uploaded already.
func cleanTemplates(files []string, records []Record) []Record {
	var c []Record

	for i := 0; i < len(records); i++ {
		record := records[i]

		if !(contains(files, record)) {
			c = append(c, record)
		}
	}

	return c
}

// Get smallest between 2 integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Generate description from name, max length 999
func generateDesc(name string) []string {
	var final []string = SeoWords
	tokens := strings.Split(name, " ")

	for i := 0; i < len(tokens); i++ {
		final = append(final, tokens[i])
	}

	Shuffle(final)

	return strings.Split(firstN(strings.Join(final, " "), 999), " ")
}

// Generate random boundary
func randomBoundary() string {
	return fmt.Sprintf(`--WebKitFormBoundary%v`, srand(16))
}

// Change order of the elements in an array
func Shuffle(slice []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for n := len(slice); n > 0; n-- {
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	}
}

// Get first 'n' characters
func firstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}
