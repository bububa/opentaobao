package drugtrace

// AlibabaalihealthdrugkytdestbilllistResult 结构体
type AlibabaalihealthdrugkytdestbilllistResult struct {
	// 接口返回model
	ModelList []AlibabaalihealthdrugkytdestbilllistModel `json:"model_list,omitempty" xml:"model_list>alibabaalihealthdrugkytdestbilllist_model,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
