package drugtrace

// AlibabaalihealthdrugkytdrdrugrecalResult 结构体
type AlibabaalihealthdrugkytdrdrugrecalResult struct {
	// 召回信息列表
	ModelList []AlibabaalihealthdrugkytdrdrugrecalModel `json:"model_list,omitempty" xml:"model_list>alibabaalihealthdrugkytdrdrugrecal_model,omitempty"`
	// 服务返回编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 服务返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
