package alitripmerchant

// AlitripMerchantGalaxyDerbyMemberVoucherCardActiveResponse 结构体
type AlitripMerchantGalaxyDerbyMemberVoucherCardActiveResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 内容
	Content *DerbyVoucherCardActiveVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
