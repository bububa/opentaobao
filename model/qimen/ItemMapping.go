package qimen

// ItemMapping 结构体
type ItemMapping struct {
	// 奇门仓储字段,C123,string(50),
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段,C123,string(50),
	ShopNick string `json:"shopNick,omitempty" xml:"shopNick,omitempty"`
	// 奇门仓储字段,C123,string(50),
	ItemSource string `json:"itemSource,omitempty" xml:"itemSource,omitempty"`
	// 奇门仓储字段,C123,string(50),
	ItemId string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// 奇门仓储字段,C123,string(50),
	ShopItemId string `json:"shopItemId,omitempty" xml:"shopItemId,omitempty"`
	// 奇门仓储字段,C123,string(50),
	SkuId string `json:"skuId,omitempty" xml:"skuId,omitempty"`
}
