package wdk

// PosRefundCreateResult 结构体
type PosRefundCreateResult struct {
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// returnCode
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// returnMsg
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
}
