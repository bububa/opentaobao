package moscm

// SalesOrderDto 结构体
type SalesOrderDto struct {
	// 订单商品明细
	OrderItems []SalesOrderItemDto `json:"order_items,omitempty" xml:"order_items>sales_order_item_dto,omitempty"`
	// 门店名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 销售类型：STORE（门店销售)、ONLINE(线上销售)
	SaleType string `json:"sale_type,omitempty" xml:"sale_type,omitempty"`
	// 最后更新时间
	LastUpdated string `json:"last_updated,omitempty" xml:"last_updated,omitempty"`
	// 订单创建时间
	DateCreated string `json:"date_created,omitempty" xml:"date_created,omitempty"`
	// 支付时间
	PaidTime string `json:"paid_time,omitempty" xml:"paid_time,omitempty"`
	// 运费，单位:元
	Freight string `json:"freight,omitempty" xml:"freight,omitempty"`
	// 订单商品金额，单位:元
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 订单号
	OrderNumber string `json:"order_number,omitempty" xml:"order_number,omitempty"`
	// 专柜ID
	CounterId string `json:"counter_id,omitempty" xml:"counter_id,omitempty"`
	// 专柜名称
	CounterName string `json:"counter_name,omitempty" xml:"counter_name,omitempty"`
	// 订单状态，可选值：UNPAID(未支付)、PAID(已支付)、PARTDISTRIBUTION(部分发货)、ALLDISTRIBUTION(全部发货)、CANCEL(取消)、BUYER_REFUND_GOODS（买家退货）
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 买家备注
	BuyerMemo string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
	// 支付信息
	Payments string `json:"payments,omitempty" xml:"payments,omitempty"`
	// 供应商专柜Id
	OutCounterId string `json:"out_counter_id,omitempty" xml:"out_counter_id,omitempty"`
	// 发货方:INTIME(银泰)、THIRD(第三方)
	Shipper string `json:"shipper,omitempty" xml:"shipper,omitempty"`
	// 配送方式:TOSTORE(门店自提)、APPOINTTIME(定时达)、EXPRESS(普通快递)、NEXTDAY(次日定时达)、TOSTORESERVICE(到店服务)、OFFLINESALE(现场销售)、NOT_REQUIRED（无需配送（虚拟商品））
	DeliveryType string `json:"delivery_type,omitempty" xml:"delivery_type,omitempty"`
	// 收货人姓名
	ReceiverName string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
	// 收货人详细地址
	ReceiverAddress string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
	// 收货人省份
	ReceiverState string `json:"receiver_state,omitempty" xml:"receiver_state,omitempty"`
	// 收货地址邮编
	ReceiverZip string `json:"receiver_zip,omitempty" xml:"receiver_zip,omitempty"`
	// 收货人移动电话
	ReceiverMobile string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
	// 收货人固话
	ReceiverPhone string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
	// 预约送达时间范围，开始时间
	AppointmentStartTime string `json:"appointment_start_time,omitempty" xml:"appointment_start_time,omitempty"`
	// 收货人所在区域
	ReceiverArea string `json:"receiver_area,omitempty" xml:"receiver_area,omitempty"`
	// 收货人所在市
	ReceiverCity string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
	// 预约送达时间范围，结束时间
	AppointmentEndTime string `json:"appointment_end_time,omitempty" xml:"appointment_end_time,omitempty"`
	// 预售订单，最晚发货时间
	LatestDeliveryTime string `json:"latest_delivery_time,omitempty" xml:"latest_delivery_time,omitempty"`
	// 订单标签，PRE_ORDER(预购)，PRE_SALE（预售），DUTCH_ACTION（东东抢）
	Tag string `json:"tag,omitempty" xml:"tag,omitempty"`
	// 发票信息（已废弃）
	Invoice *InvoiceDto `json:"invoice,omitempty" xml:"invoice,omitempty"`
}
