package alitripmerchant

// AlitripmerchantgalaxycommongetenumsbynameResponse 结构体
type AlitripmerchantgalaxycommongetenumsbynameResponse struct {
	// 枚举
	Content []EnumVo `json:"content,omitempty" xml:"content>enum_vo,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
