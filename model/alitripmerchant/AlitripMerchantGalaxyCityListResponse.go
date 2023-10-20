package alitripmerchant

// AlitripmerchantgalaxycitylistResponse 结构体
type AlitripmerchantgalaxycitylistResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回城市列表
	Content *AddressListSearchDto `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
