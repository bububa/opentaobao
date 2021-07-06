package aliospay

// GetTokenRequest 结构体
type GetTokenRequest struct {
	// 请求唯一id，不可重复，服务端会根据此参数防重放
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 语言,en表示英文，zh表示中文
	Lang string `json:"lang,omitempty" xml:"lang,omitempty"`
	// 请求时间戳
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 业务订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 订单标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// ORDER订单token，OTHER其他token
	TokenType string `json:"token_type,omitempty" xml:"token_type,omitempty"`
	// 订单总金额
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 参与优惠计算的金额，用此字段用于让订单中部分金额不参与优惠的计算
	DiscountableAmount int64 `json:"discountable_amount,omitempty" xml:"discountable_amount,omitempty"`
}
