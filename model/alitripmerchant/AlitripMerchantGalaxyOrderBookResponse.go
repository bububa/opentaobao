package alitripmerchant

// AlitripMerchantGalaxyOrderBookResponse 结构体
type AlitripMerchantGalaxyOrderBookResponse struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 预订返回
	Content *PayParamResult `json:"content,omitempty" xml:"content,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
