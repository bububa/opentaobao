package baoxian

// AlipayBaoxianClaimUpdateMtopResult 结构体
type AlipayBaoxianClaimUpdateMtopResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
