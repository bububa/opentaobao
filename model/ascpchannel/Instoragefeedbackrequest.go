package ascpchannel

// Instoragefeedbackrequest 结构体
type Instoragefeedbackrequest struct {
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 逆向履约单号
	BizOrderCode string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
	// ERP业务编码
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 逆向入库时间
	InstorageTime string `json:"instorage_time,omitempty" xml:"instorage_time,omitempty"`
	// 物流公司编码
	TmsServiceCode string `json:"tms_service_code,omitempty" xml:"tms_service_code,omitempty"`
	// 快递单号
	TmsOrderCode string `json:"tms_order_code,omitempty" xml:"tms_order_code,omitempty"`
	// 退回收件人信息(商家)
	ReceiverInfo *Receiverinfo `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
	// 退回寄件人信息(消费者)
	SenderInfo *Senderinfo `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
	// 退回订单货品信息列表
	OrderItems []Orderitems `json:"order_items,omitempty" xml:"order_items>orderitems,omitempty"`
	// 退回仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
}
