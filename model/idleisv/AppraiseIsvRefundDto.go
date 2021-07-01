package idleisv

// AppraiseIsvRefundDto 结构体
type AppraiseIsvRefundDto struct {
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 商品购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 买家申请退款原因
	BuyerApplyReason string `json:"buyer_apply_reason,omitempty" xml:"buyer_apply_reason,omitempty"`
	// 买家退款说明,买家申请退款二级原因
	BuyerApplySubReason string `json:"buyer_apply_sub_reason,omitempty" xml:"buyer_apply_sub_reason,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 货物状态
	GoodsStatus int64 `json:"goods_status,omitempty" xml:"goods_status,omitempty"`
	// 商品信息
	Item *AppraiseIsvItemDto `json:"item,omitempty" xml:"item,omitempty"`
	// 买家是否需要退货
	NeedReturnGoods bool `json:"need_return_goods,omitempty" xml:"need_return_goods,omitempty"`
	// 订单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单实付金额,单位分
	Payment int64 `json:"payment,omitempty" xml:"payment,omitempty"`
	// 退款申请时间,时间戳,单位分
	RefundCreateTime int64 `json:"refund_create_time,omitempty" xml:"refund_create_time,omitempty"`
	// 退还金额(退还给买家的金额),单位分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退款订单号
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 退货-快递公司
	RefundPostCompany string `json:"refund_post_company,omitempty" xml:"refund_post_company,omitempty"`
	// 退货-快递单号
	RefundPostNo string `json:"refund_post_no,omitempty" xml:"refund_post_no,omitempty"`
	// 退款订单状态
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 卖家同意退货说明
	SellerAgreeMsg string `json:"seller_agree_msg,omitempty" xml:"seller_agree_msg,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 卖家拒绝退款原因
	SellerRefuseReason string `json:"seller_refuse_reason,omitempty" xml:"seller_refuse_reason,omitempty"`
	// 买家申请退款描述
	BuyerApplyDesc string `json:"buyer_apply_desc,omitempty" xml:"buyer_apply_desc,omitempty"`
	// 退款最新修改时间,时间戳,单位分
	RefundModifyTime int64 `json:"refund_modify_time,omitempty" xml:"refund_modify_time,omitempty"`
	// 卖家拒绝退款说明
	SellerRefuseMsg string `json:"seller_refuse_msg,omitempty" xml:"seller_refuse_msg,omitempty"`
	// 退款超时信息
	TimeoutData *IsvRefundTimeoutDto `json:"timeout_data,omitempty" xml:"timeout_data,omitempty"`
}
