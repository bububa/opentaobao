package alitripmerchant

// AlitripMerchantGalaxyCommonGetEnumsbynameResponse 结构体
type AlitripMerchantGalaxyCommonGetEnumsbynameResponse struct {
	// 枚举
	Content []EnumVO `json:"content,omitempty" xml:"content>enum_vo,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
