package handlers

import (
	"fmt"
	"rocat/api"

	"github.com/urfave/cli/v2"
)

func WhoAmI(cCtx *cli.Context, cookie string) error {
	user, err := api.GetUserInfo(cookie)

	if err != nil {
		fmt.Println(`Unable to fetch user info, please re-check your cookie.`)
	}

	if user.UserId > 0 {
		fmt.Println(fmt.Sprintf("Username: %v\nUserId: %v\nRobux: %v\nHas Premium: %v\nHas Builders Club: %v",
			user.UserName,
			user.UserId,
			user.RobuxBalance,
			user.IsPremium,
			user.IsAnyBuildersClubMember,
		))
	}

	return nil
}
