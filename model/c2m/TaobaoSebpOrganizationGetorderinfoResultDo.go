package c2m

// TaobaoSebpOrganizationGetorderinfoResultDo 结构体
type TaobaoSebpOrganizationGetorderinfoResultDo struct {
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 结果信息
	Module *PageInfo `json:"module,omitempty" xml:"module,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
