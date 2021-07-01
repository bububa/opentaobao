package idleisv

// AlibabaIdleIsvOrderQueryResult 结构体
type AlibabaIdleIsvOrderQueryResult struct {
	// 订单信息
	Module *AppraiseIsvOrderDto `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
