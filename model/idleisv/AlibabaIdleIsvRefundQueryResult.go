package idleisv

// AlibabaidleisvrefundqueryResult 结构体
type AlibabaidleisvrefundqueryResult struct {
	// 退款信息
	Module *AppraiseIsvRefundDto `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
