package alitripmerchant

// AlitripMerchantGalaxyMemberQueryResponse 结构体
type AlitripMerchantGalaxyMemberQueryResponse struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 用户信息
	Content *MemberDto `json:"content,omitempty" xml:"content,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
