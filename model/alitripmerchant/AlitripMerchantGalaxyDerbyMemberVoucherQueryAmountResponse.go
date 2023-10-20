package alitripmerchant

// AlitripmerchantgalaxyderbymembervoucherqueryamountResponse 结构体
type AlitripmerchantgalaxyderbymembervoucherqueryamountResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 实体类
	Content *DerbyMemberVo `json:"content,omitempty" xml:"content,omitempty"`
	// 成功/失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
