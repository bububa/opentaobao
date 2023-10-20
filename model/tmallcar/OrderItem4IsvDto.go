package tmallcar

// OrderItem4isvDto 结构体
type OrderItem4isvDto struct {
	// 实际支付金额
	ActualTotalFee string `json:"actual_total_fee,omitempty" xml:"actual_total_fee,omitempty"`
	// 扩展属性
	Extension string `json:"extension,omitempty" xml:"extension,omitempty"`
	// 商品头图
	ItemHeadImg string `json:"item_head_img,omitempty" xml:"item_head_img,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 订单商品原价
	TotalFee string `json:"total_fee,omitempty" xml:"total_fee,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 是否在三包期内
	InGuarantee bool `json:"in_guarantee,omitempty" xml:"in_guarantee,omitempty"`
}
