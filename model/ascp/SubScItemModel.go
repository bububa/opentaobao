package ascp

// SubScItemModel 结构体
type SubScItemModel struct {
	// 货品商家编码
	ScitemCode string `json:"scitem_code,omitempty" xml:"scitem_code,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 子货品id
	ScitemId string `json:"scitem_id,omitempty" xml:"scitem_id,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 子货品单价(人命币则为分)
	RetailPrice int64 `json:"retail_price,omitempty" xml:"retail_price,omitempty"`
	// 是否固定价格 1是 0否
	FixedPrice int64 `json:"fixed_price,omitempty" xml:"fixed_price,omitempty"`
}
