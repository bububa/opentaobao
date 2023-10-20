package shop

// AlibabaShopCouponApplyResult 结构体
type AlibabaShopCouponApplyResult struct {
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 操作成功或者失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
