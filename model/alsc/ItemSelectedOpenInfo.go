package alsc

// ItemSelectedOpenInfo 结构体
type ItemSelectedOpenInfo struct {
	// 规则ID
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 菜品ID
	DishId string `json:"dish_id,omitempty" xml:"dish_id,omitempty"`
	// 外部菜品ID
	DishOutNo string `json:"dish_out_no,omitempty" xml:"dish_out_no,omitempty"`
	// 外部规则ID
	SkuOutNo string `json:"sku_out_no,omitempty" xml:"sku_out_no,omitempty"`
}
