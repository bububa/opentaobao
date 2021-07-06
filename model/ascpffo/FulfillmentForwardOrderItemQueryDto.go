package ascpffo

// FulfillmentForwardOrderItemQueryDto 结构体
type FulfillmentForwardOrderItemQueryDto struct {
	// 履约单号
	FulfillmentOrderNo string `json:"fulfillment_order_no,omitempty" xml:"fulfillment_order_no,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}
