package servicecenter

// ResultBase 结构体
type ResultBase struct {
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// value
	Value *TpFundsRecoverResultDo `json:"value,omitempty" xml:"value,omitempty"`
	// errorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 查询接口是否OK
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
