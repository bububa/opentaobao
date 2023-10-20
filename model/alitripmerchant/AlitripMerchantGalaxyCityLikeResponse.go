package alitripmerchant

// AlitripmerchantgalaxycitylikeResponse 结构体
type AlitripmerchantgalaxycitylikeResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回结果
	Content *AddressSearchVo `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
