package promotion

// MarketResult 结构体
type MarketResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回的匿名码对象
	Data *CouponActivity `json:"data,omitempty" xml:"data,omitempty"`
}
