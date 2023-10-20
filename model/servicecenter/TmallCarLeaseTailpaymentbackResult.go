package servicecenter

// TmallcarleasetailpaymentbackResult 结构体
type TmallcarleasetailpaymentbackResult struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// costTime
	CostTime int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// gmtCurrentTime
	GmtCurrentTime int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
