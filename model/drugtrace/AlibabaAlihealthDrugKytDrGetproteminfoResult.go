package drugtrace

// AlibabaalihealthdrugkytdrgetproteminfoResult 结构体
type AlibabaalihealthdrugkytdrgetproteminfoResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *AlibabaalihealthdrugkytdrgetproteminfoModel `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
