package drugtrace

// AlibabaalihealthdrugkytscqyputpackageunbindResultModel 结构体
type AlibabaalihealthdrugkytscqyputpackageunbindResultModel struct {
	// 消息码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 消息提示内容
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 内层大对象
	Model *CodeMovePackagingSecondDto `json:"model,omitempty" xml:"model,omitempty"`
	// 查询成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
