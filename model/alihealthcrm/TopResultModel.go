package alihealthcrm

// TopResultModel 结构体
type TopResultModel struct {
	// 调用结果
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 调用返回code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
