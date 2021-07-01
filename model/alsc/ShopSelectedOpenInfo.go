package alsc

// ShopSelectedOpenInfo 结构体
type ShopSelectedOpenInfo struct {
	// 门店ID
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部门店ID
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
}
