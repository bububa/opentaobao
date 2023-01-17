package tmallsc

// ServicePriceQueryRequest 结构体
type ServicePriceQueryRequest struct {
	// 服务外部编码
	ServiceOuterIds []string `json:"service_outer_ids,omitempty" xml:"service_outer_ids>string,omitempty"`
	// 卖家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
