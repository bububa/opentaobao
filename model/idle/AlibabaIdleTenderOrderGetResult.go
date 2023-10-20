package idle

// AlibabaidletenderordergetResult 结构体
type AlibabaidletenderordergetResult struct {
	// 错误code
	ErrorCodeInfo string `json:"error_code_info,omitempty" xml:"error_code_info,omitempty"`
	// 错误msg
	ErrorMsgInfo string `json:"error_msg_info,omitempty" xml:"error_msg_info,omitempty"`
	// 结果详细信息
	Module *TenderOrderInfoDto `json:"module,omitempty" xml:"module,omitempty"`
	// 是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
}
