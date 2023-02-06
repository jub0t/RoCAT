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
