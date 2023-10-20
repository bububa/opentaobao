package drugtrace

// AlibabaalihealthdrugkytdrbillcheckResult 结构体
type AlibabaalihealthdrugkytdrbillcheckResult struct {
	// 服务返回码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 服务返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *AlibabaalihealthdrugkytdrbillcheckModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
