package omniorder

// TradeItemInfoDto 结构体
type TradeItemInfoDto struct {
	// 商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 商品单价，单位为分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 商品数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 体积，单位为升
	Volumn int64 `json:"volumn,omitempty" xml:"volumn,omitempty"`
	// 重量，单位为克
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
}
