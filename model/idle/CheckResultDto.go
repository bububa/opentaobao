package idle

// CheckResultDto 结构体
type CheckResultDto struct {
	// 不通过错误码
	ExtraCode string `json:"extra_code,omitempty" xml:"extra_code,omitempty"`
	// 不通过详细信息
	ExtraMessage string `json:"extra_message,omitempty" xml:"extra_message,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// sku id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 是否通过
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
