package fenxiao

// BuyerRefund 
type BuyerRefund struct {

    // 退款创建时间
    RefundCreateTime   string `json:"refund_create_time,omitempty"`

    // 退款状态
    RefundStatus   int64 `json:"refund_status,omitempty"`

    // 货物的状态：
买家已收到货
买家已退货
买家未收到货
    GoodsStatusDesc   string `json:"goods_status_desc,omitempty"`

    // 买家是否退货
    NeedReturnGoods   bool `json:"need_return_goods,omitempty"`

    // 退还买家的金额
    ReturnFee   int64 `json:"return_fee,omitempty"`

    // 支付分销商的金额
    ToSellerFee   int64 `json:"to_seller_fee,omitempty"`

    // 退款原因
    RefundReason   string `json:"refund_reason,omitempty"`

    // 退款说明
    RefundDesc   string `json:"refund_desc,omitempty"`

    // 交易退款id
    RefundId   int64 `json:"refund_id,omitempty"`

    // 采购单子单id
    SubOrderId   int64 `json:"sub_order_id,omitempty"`

    // 订单id
    BizOrderId   int64 `json:"biz_order_id,omitempty"`

    // 下游买家nick
    BuyerNick   string `json:"buyer_nick,omitempty"`

    // 退款修改时间。格式:yyyy-MM-dd HH:mm:ss
    Modified   string `json:"modified,omitempty"`

}
