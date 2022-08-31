package btrip

// AirlineInfo 结构体
type AirlineInfo struct {
	// 航司编码
	AirlineCode string `json:"airline_code,omitempty" xml:"airline_code,omitempty"`
	// 航司名称
	AirlineName string `json:"airline_name,omitempty" xml:"airline_name,omitempty"`
	// 航司简称
	AirlineSimpleName string `json:"airline_simple_name,omitempty" xml:"airline_simple_name,omitempty"`
}
