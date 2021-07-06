package ascpchannel

// PresalesorderTest 结构体
type PresalesorderTest struct {
	// 订单信息
	OrderLines []Orderlines `json:"order_lines,omitempty" xml:"order_lines>orderlines,omitempty"`
	// 出库单号
	PresalesOrderCode string `json:"presales_order_code,omitempty" xml:"presales_order_code,omitempty"`
	// 前台订单(店铺订单)创建时间(下单时间)
	PlaceOrderTime string `json:"place_order_time,omitempty" xml:"place_order_time,omitempty"`
	// 订单总金额=应收金额+已收金额=商品总金额-订单折扣金额+快递费用
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 仓库
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 发件人信息
	SenderInfo *Senderinfo `json:"sender_info,omitempty" xml:"sender_info,omitempty"`
	// 收件人信息
	ReceiverInfo *Receiverinfo `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
}
