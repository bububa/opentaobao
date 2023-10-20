package tmallcarenter

// TmallCarcenterVehicleVersionInsertResult 结构体
type TmallCarcenterVehicleVersionInsertResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// object
	Object string `json:"object,omitempty" xml:"object,omitempty"`
	// gmtCurrentTime
	GmtCurrentTime int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
	// costTime
	CostTime int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
