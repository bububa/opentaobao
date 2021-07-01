package tmallservice

// TmallServicecenterWorkcardSuspendResult 结构体
type TmallServicecenterWorkcardSuspendResult struct {
	// 是否失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误原因描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
}
