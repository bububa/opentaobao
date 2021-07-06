package aliospay

// RefundRequest 结构体
type RefundRequest struct {
	// 请求唯一id，不可重复，服务端会根据此参数防重放
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 语言,en表示英文，zh表示中文
	Lang string `json:"lang,omitempty" xml:"lang,omitempty"`
	// 请求时间戳
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 业务订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 标识一次退款请求，保证唯一
	OutRequestNo string `json:"out_request_no,omitempty" xml:"out_request_no,omitempty"`
	// 退款金额，单位分
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}
