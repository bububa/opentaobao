package fenxiao

// BuyerRefund 结构体
type BuyerRefund struct {
	// 消费者订单退款创建时间
	RefundCreateTime string `json:"refund_create_time,omitempty" xml:"refund_create_time,omitempty"`
	// 货物的状态：买家已收到货买家已退货买家未收到货
	GoodsStatusDesc string `json:"goods_status_desc,omitempty" xml:"goods_status_desc,omitempty"`
	// 消费者退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 消费者退款说明
	RefundDesc string `json:"refund_desc,omitempty" xml:"refund_desc,omitempty"`
	// 消费者nick
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 消费者退款修改时间。格式:yyyy-MM-dd HH:mm:ss
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
	// 消费者淘宝id的加密key
	OpenBuyerId string `json:"open_buyer_id,omitempty" xml:"open_buyer_id,omitempty"`
	// 消费者退款状态最后修改时间，格式 yyyy-MM-dd HH:mm:ss
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 分销子订单号
	SubOrderId int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 消费者订单对应的退款单号
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 消费者订单退款涉及的消费者正向子订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 消费者订单退款状态 1、消费者已经申请退款，等待分销商确认 2、分销商已经同意退货，等待消费者退货  3、消费者已经退货，等待分销商确认收货 4、退款关闭   5、退款成功 6、分销商拒绝退款,待消费者重新修改  7、等待消费者确认重新邮寄的货物
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 退还给消费者的金额(分)
	ReturnFee int64 `json:"return_fee,omitempty" xml:"return_fee,omitempty"`
	// 确认收货后会打款给分销商的金额(分),分摊到子单的实付金额-退款给消费者的金额
	ToSellerFee int64 `json:"to_seller_fee,omitempty" xml:"to_seller_fee,omitempty"`
	// 消费者退货数量
	ReturnGoodsQuantity int64 `json:"return_goods_quantity,omitempty" xml:"return_goods_quantity,omitempty"`
	// 买家是否退货
	NeedReturnGoods bool `json:"need_return_goods,omitempty" xml:"need_return_goods,omitempty"`
}
