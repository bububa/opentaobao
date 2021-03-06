package tmallnr

// TmallNrtNewcouponSendResult 结构体
type TmallNrtNewcouponSendResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 发券结果对象
	Model *SendCouponResponse `json:"model,omitempty" xml:"model,omitempty"`
}
