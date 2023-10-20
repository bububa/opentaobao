package drugtrace

// AlibabaalihealthdrugkytdrgetupteminfoResult 结构体
type AlibabaalihealthdrugkytdrgetupteminfoResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	Model *AlibabaalihealthdrugkytdrgetupteminfoModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
