package alihouse

// BrokerPointResultDto 结构体
type BrokerPointResultDto struct {
	// 1
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 1
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
