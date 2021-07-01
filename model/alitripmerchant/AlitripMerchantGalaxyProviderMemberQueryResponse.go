package alitripmerchant

// AlitripMerchantGalaxyProviderMemberQueryResponse 结构体
type AlitripMerchantGalaxyProviderMemberQueryResponse struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回结果
	ProviderMemberVo *ProviderMemberVo `json:"provider_member_vo,omitempty" xml:"provider_member_vo,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
