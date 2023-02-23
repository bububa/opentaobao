package opentrade

// AbilityResponse 结构体
type AbilityResponse struct {
	// 保存信息的参数
	PriceKey string `json:"price_key,omitempty" xml:"price_key,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// true or false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
