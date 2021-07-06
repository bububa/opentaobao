package tmallservice

// TmallServicecenterWorkcardReservefailResult 结构体
type TmallServicecenterWorkcardReservefailResult struct {
	// 用于对外展示的错误信息
	DisplayMsg string `json:"display_msg,omitempty" xml:"display_msg,omitempty"`
	// 消息代码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 系统级错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
