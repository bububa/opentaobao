package tmallsc

// ServicePriceQueryResponse 结构体
type ServicePriceQueryResponse struct {
	// 服务价格记录行
	ServicePriceLines []ServicePriceLine `json:"service_price_lines,omitempty" xml:"service_price_lines>service_price_line,omitempty"`
}
