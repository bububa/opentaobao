package tmallnr

// TmallNrtCouponInstanceQueryResult 结构体
type TmallNrtCouponInstanceQueryResult struct {
	// 券实例实体
	Models []string `json:"models,omitempty" xml:"models>string,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
