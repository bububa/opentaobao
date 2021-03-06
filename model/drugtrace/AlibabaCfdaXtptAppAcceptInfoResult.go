package drugtrace

// AlibabaCfdaXtptAppAcceptInfoResult 结构体
type AlibabaCfdaXtptAppAcceptInfoResult struct {
	// 消息编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 结果信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
