package idleisv

// AlibabaIdleIsvRefundQueryResult 
type AlibabaIdleIsvRefundQueryResult struct {
    // 退款信息
    Module   *AppraiseIsvRefundDTO `json:"module,omitempty" xml:"module,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
