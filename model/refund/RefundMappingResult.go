package refund

// RefundMappingResult 结构体
type RefundMappingResult struct {
	// 结果信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 退款ID
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}
