package promotion

// AlibababenefitdrawResult 结构体
type AlibababenefitdrawResult struct {
	// message
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}
