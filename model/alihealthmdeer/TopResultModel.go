package alihealthmdeer

// TopResultModel 结构体
type TopResultModel struct {
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
