package main

type ResponseItems struct {
	Id                      int      `json:"id"`
	ItemType                string   `json:"itemType"`
	AssetType               int      `json:"assetType"`
	Name                    string   `json:"name"`
	Description             string   `json:"description"`
	ProductId               int      `json:"productId"`
	Genres                  []string `json:"genres"`
	ItemStatus              []string `json:"itemStatus"`
	ItemRestrictions        []string `json:"itemRestrictions"`
	CreatorHasVerifiedBadge bool     `json:"creatorHasVerifiedBadge"`
	CreatorType             string   `json:"creatorType"`
	CreatorTargetId         int      `json:"creatorTargetId"`
	CreatorName             string   `json:"creatorName"`
	Price                   int      `json:"price"`
	PriceStatus             string   `json:"priceStatus"`
	OffSaleDeadline         string   `json:"offSaleDeadline"`
	IsNew                   bool     `json:"isNew"`
	IsLimited               bool     `json:"isLimited"`
	IsLimitedUnique         bool     `json:"isLimitedUnique"`
	Remaining               int      `json:"remaining"`
	MinimumMembershipLevel  int      `json:"minimumMembershipLevel"`
}

type RequestItems struct {
	Id            int    `json:"id"`
	Key           string `json:"key"`
	ItemType      string `json:"itemType"`
	ThumbnailType string `json:"thumbnailType"`
}

type CatalogueItem struct {
	Id       int    `json:"id"`
	ItemType string `json:"itemType"`
}

type CatalogueResponse struct {
	NextPageCursor string          `json:"nextPageCursor"`
	Data           []CatalogueItem `json:"data"`
}

type GetClothesRequest struct {
	Items []CatalogueItem `json:"items"`
}

type GetClothesResponse struct {
	Data []ResponseItems `json:"data"`
}

type AssetInfo struct {
	Location             string `json:"location"`
	RequestId            string `json:"requestId"`
	IsArchived           bool   `json:"isArchived"`
	IsCopyrightProtected bool   `json:"IsCopyrightProtected"`
}

type Record struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type AccountBalanceResponse struct {
	Robux int `json:"robux"`
}

type UserInfo struct {
	UserId                  int    `json:"UserID"`
	UserName                string `json:"UserName"`
	RobuxBalance            int    `json:"RobuxBalance"`
	ThumbnailUrl            string `json:"ThumbnailUrl"`
	IsAnyBuildersClubMember bool   `json:"IsAnyBuildersClubMember"`
	IsPremium               bool   `json:"IsPremium"`
}

type UploadConfig struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	CreatorTargetId string `json:"creatorTargetId"`
	CreatorType     string `json:"creatorType"`
}
