package aliospay

// QueryTradeRequest 
type QueryTradeRequest struct {
    // 请求唯一id，不可重复，服务端会根据此参数防重放
    TraceId   string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
    // 语言,en表示英文，zh表示中文
    Lang   string `json:"lang,omitempty" xml:"lang,omitempty"`
    // 发送请求的时间戳
    Time   string `json:"time,omitempty" xml:"time,omitempty"`
    // 业务订单号
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    // alios支付订单id
    PayOrderId   string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
}
