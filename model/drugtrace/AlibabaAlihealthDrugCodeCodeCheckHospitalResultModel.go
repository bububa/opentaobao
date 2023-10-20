package drugtrace

// AlibabaalihealthdrugcodecodecheckhospitalResultModel 结构体
type AlibabaalihealthdrugcodecodecheckhospitalResultModel struct {
	// 成功失败编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 成功失败描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 成功失败结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 成功失败标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
