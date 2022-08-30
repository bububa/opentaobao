package tmallcar

// CreditLoanStatusSyncResp 结构体
type CreditLoanStatusSyncResp struct {
	// 是否成功
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
	// 是否重试
	Retry bool `json:"retry,omitempty" xml:"retry,omitempty"`
}
