package drugtrace

// AlibabaalihealthdruguploadextinfoResultModel 结构体
type AlibabaalihealthdruguploadextinfoResultModel struct {
	// 编码值
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 中文解释
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口响应是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 业务处理是否成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}
