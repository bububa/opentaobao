package alihealth2

// NormalGoodsVo 结构体
type NormalGoodsVo struct {
	// 标品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 成本价，单位元
	CostPrice string `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
	// 售价，单位元
	SoldPrice string `json:"sold_price,omitempty" xml:"sold_price,omitempty"`
	// 标品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}
