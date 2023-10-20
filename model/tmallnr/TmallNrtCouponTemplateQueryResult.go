package tmallnr

// TmallnrtcoupontemplatequeryResult 结构体
type TmallnrtcoupontemplatequeryResult struct {
	// 券模版数据
	Model []NrtCouponTemplateDto `json:"model,omitempty" xml:"model>nrt_coupon_template_dto,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}
