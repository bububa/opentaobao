package alihouse

// TmallAlihouseTradeCouponRefundOrderQueryResult 结构体
type TmallAlihouseTradeCouponRefundOrderQueryResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 结果说明
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 返回素材id
	Data *CouponRefundOrderStatusDto `json:"data,omitempty" xml:"data,omitempty"`
}
