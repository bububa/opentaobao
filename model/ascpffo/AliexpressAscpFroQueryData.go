package ascpffo

// AliexpressAscpFroQueryData 结构体
type AliexpressAscpFroQueryData struct {
	// 用户订单号
	TradeOrderNo string `json:"trade_order_no,omitempty" xml:"trade_order_no,omitempty"`
	// 履约单号
	FulfillmentOrderNo string `json:"fulfillment_order_no,omitempty" xml:"fulfillment_order_no,omitempty"`
	// 原物流单号
	OriginalLbxNo string `json:"original_lbx_no,omitempty" xml:"original_lbx_no,omitempty"`
	// 物流单号
	LbxNo string `json:"lbx_no,omitempty" xml:"lbx_no,omitempty"`
	// 收货人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 收货人手机
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 物流公司
	ShippingProviderName string `json:"shipping_provider_name,omitempty" xml:"shipping_provider_name,omitempty"`
	// 运单号
	TrackingNumber string `json:"tracking_number,omitempty" xml:"tracking_number,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 退货类型
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 扩展字段
	ExtendFields string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
	// 退货下单时间戳
	ReturnOrderCreateTime int64 `json:"return_order_create_time,omitempty" xml:"return_order_create_time,omitempty"`
	// 退货入库时间戳
	ReturnOrderInboundTime int64 `json:"return_order_inbound_time,omitempty" xml:"return_order_inbound_time,omitempty"`
}
