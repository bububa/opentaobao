package alihouse

// TmallalihousetradecouponorderstatusqueryResult 结构体
type TmallalihousetradecouponorderstatusqueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果说明
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回素材id
	Data *CouponOrderStatusDto `json:"data,omitempty" xml:"data,omitempty"`
}
