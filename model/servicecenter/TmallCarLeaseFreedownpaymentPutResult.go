package servicecenter

// TmallCarLeaseFreedownpaymentPutResult 结构体
type TmallCarLeaseFreedownpaymentPutResult struct {
	// 当前时间
	GmtCurrentTime int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回对象,成功：true，失败：false
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 耗费时间
	CostTime int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
}
