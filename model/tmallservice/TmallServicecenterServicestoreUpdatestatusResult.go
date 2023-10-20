package tmallservice

// TmallservicecenterservicestoreupdatestatusResult 结构体
type TmallservicecenterservicestoreupdatestatusResult struct {
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
