package idleisv

// AlibabaIdleIsvOrderQueryResult 
type AlibabaIdleIsvOrderQueryResult struct {
    // 订单信息
    Module   *AppraiseIsvOrderDTO `json:"module,omitempty" xml:"module,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
