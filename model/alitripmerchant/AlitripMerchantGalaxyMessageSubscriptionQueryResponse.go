package alitripmerchant

// AlitripMerchantGalaxyMessageSubscriptionQueryResponse 结构体
type AlitripMerchantGalaxyMessageSubscriptionQueryResponse struct {
	// 错误编码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误码
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 权限是否存在
	Content bool `json:"content,omitempty" xml:"content,omitempty"`
}
