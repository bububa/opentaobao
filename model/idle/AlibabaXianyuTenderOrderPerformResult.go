package idle

// AlibabaXianyuTenderOrderPerformResult 结构体
type AlibabaXianyuTenderOrderPerformResult struct {
	// 错误信息
	ErrCodeInfo string `json:"err_code_info,omitempty" xml:"err_code_info,omitempty"`
	// 错误码
	ErrMsgInfo string `json:"err_msg_info,omitempty" xml:"err_msg_info,omitempty"`
	// 请求是否成功
	RespSuccess bool `json:"resp_success,omitempty" xml:"resp_success,omitempty"`
	// 结果是否成功
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}
