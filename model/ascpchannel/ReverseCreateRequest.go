package ascpchannel

// ReverseCreateRequest 结构体
type ReverseCreateRequest struct {
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// ERP业务编码
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 逆向类型:1=客退;2=运配异常;3=拒签退回;4=拦截退回;5=上门取退
	ReverseType string `json:"reverse_type,omitempty" xml:"reverse_type,omitempty"`
	// 原正向发货履约主单号
	BizOrderCode string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
	// 快递公司code.调用 taobao.logistics.companies.get 获取
	TmsServiceCode string `json:"tms_service_code,omitempty" xml:"tms_service_code,omitempty"`
	// 运单号
	TmsOrderCode string `json:"tms_order_code,omitempty" xml:"tms_order_code,omitempty"`
	// 退回收件人信息（商家）
	ReceiverInfo *Receiverinfo `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
	// 退回寄件人信息（消费者）
	SenderInfo *Senderinfo `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
	// 退回订单货品信息列表
	OrderItems []Orderitems `json:"order_items,omitempty" xml:"order_items>orderitems,omitempty"`
	// 退回仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
}
