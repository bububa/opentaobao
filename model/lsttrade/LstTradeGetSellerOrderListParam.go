package lsttrade

// LstTradeGetSellerOrderListParam 结构体
type LstTradeGetSellerOrderListParam struct {
	// 买家id
	BuyerMemberId string `json:"buyer_member_id,omitempty" xml:"buyer_member_id,omitempty"`
	// 订单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 退款状态，支持：     "waitselleragree"(等待卖家同意),     "refundsuccess"(退款成功),     "refundclose"(退款关闭),     "waitbuyermodify"(待买家修改),     "waitbuyersend"(等待买家退货),     "waitsellerreceive"(等待卖家确认收货)
	RefundStatus string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 商品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 查询下单开始时间
	CreateStartTime string `json:"create_start_time,omitempty" xml:"create_start_time,omitempty"`
	// 查询下单开始时间，时间间隔一个月内
	CreateEndTime string `json:"create_end_time,omitempty" xml:"create_end_time,omitempty"`
	// 交易订单时间
	TradeEndTime string `json:"trade_end_time,omitempty" xml:"trade_end_time,omitempty"`
	// 交易订单开始时间
	TradeStartTime string `json:"trade_start_time,omitempty" xml:"trade_start_time,omitempty"`
	// 更新时间-查询开始
	UpdateStartTime string `json:"update_start_time,omitempty" xml:"update_start_time,omitempty"`
	// 更新时间-查询截止时间
	UpdateEndTime string `json:"update_end_time,omitempty" xml:"update_end_time,omitempty"`
	// 主订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 子订单
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 当前页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 是否查询历史
	IsHis bool `json:"is_his,omitempty" xml:"is_his,omitempty"`
}
