package alsc

// ShopSelectedOpenInfoList 结构体
type ShopSelectedOpenInfoList struct {
	// 门店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部门店id
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
}
