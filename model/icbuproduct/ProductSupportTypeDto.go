package icbuproduct

// ProductSupportTypeDto 结构体
type ProductSupportTypeDto struct {
	// 是否支持下单品
	SupportPostWholeSale bool `json:"support_post_whole_sale,omitempty" xml:"support_post_whole_sale,omitempty"`
	// 是否支持询盘品
	SupportPostSourcing bool `json:"support_post_sourcing,omitempty" xml:"support_post_sourcing,omitempty"`
}
