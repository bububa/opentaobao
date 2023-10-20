package drugtrace

// AlibabaalihealthdrugkytybgetcoderelationResult 结构体
type AlibabaalihealthdrugkytybgetcoderelationResult struct {
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// model
	ModelList *AlibabaalihealthdrugkytybgetcoderelationModel `json:"model_list,omitempty" xml:"model_list,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
