package tmallsc

// TmallServicecenterCallrecordQueryResult 结构体
type TmallServicecenterCallrecordQueryResult struct {
	// 通话记录信息
	Value []ServiceCallRecordCO `json:"value,omitempty" xml:"value>service_call_record_co,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
