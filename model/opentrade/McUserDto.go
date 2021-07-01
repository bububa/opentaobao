package opentrade

// McUserDto 结构体
type McUserDto struct {
	// 用户状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 排队活动ID
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 排队传入的扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 排队商品数量
	Quality int64 `json:"quality,omitempty" xml:"quality,omitempty"`
	// 排队商品SKU ID
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 排队商品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 用户openId
	UserOpenId string `json:"user_open_id,omitempty" xml:"user_open_id,omitempty"`
}
