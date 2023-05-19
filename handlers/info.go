package handlers

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func Info(cCtx *cli.Context) error {
	dirs, err := os.ReadDir("./downloads")

	if err != nil {
		print("Error occured while getting downloads")
		return err
	}

	fmt.Println(fmt.Sprintf("Downloaded Clothing: %v", len(dirs)))

	return nil
}
