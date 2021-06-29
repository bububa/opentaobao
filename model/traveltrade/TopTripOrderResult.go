package traveltrade

// TopTripOrderResult 
type TopTripOrderResult struct {
    // 主订单id
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // string类型订单id
    OrderIdString   string `json:"order_id_string,omitempty" xml:"order_id_string,omitempty"`
    // 订单创建时间
    CreatedTime   string `json:"created_time,omitempty" xml:"created_time,omitempty"`
    // 订单修改时间
    ModifiedTime   string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
    // 交易结束时间。交易成功时间(更新交易状态为成功的同时更新)/确认收货时间或者交易关闭时间
    EndTime   string `json:"end_time,omitempty" xml:"end_time,omitempty"`
    // 主订单状态。交易状态。可选值: TRADE_NO_CREATE_PAY(没有创建支付宝交易)；WAIT_BUYER_PAY(等待买家付款)；SELLER_CONSIGNED_PART(卖家部分发货)；WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)；WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)；TRADE_FINISHED(交易成功)；TRADE_CLOSED(付款以后用户退款成功，交易自动关闭)；TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)；PAY_PENDING(国际信用卡支付付款确认中);
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 交易类型列表，fixed(一口价);auction(拍卖);guarantee_trade(一口价、拍卖);auto_delivery(自动发货);cod(货到付款);external_trade(统一外部交易);instant_trade(即时交易);b2c_cod(大商家货到付款);nopaid(即时到帐);eticket(电子凭证订单);step(分阶段付款)
    Type   string `json:"type,omitempty" xml:"type,omitempty"`
    // 主订单支付信息
    PayInfo   *PayInfo `json:"pay_info,omitempty" xml:"pay_info,omitempty"`
    // 卖家信息
    SellerInfo   *SellerInfo `json:"seller_info,omitempty" xml:"seller_info,omitempty"`
    // 买家信息
    BuyerInfo   *BuyerInfo `json:"buyer_info,omitempty" xml:"buyer_info,omitempty"`
    // 主、子订单优惠信息
    PromotionDetails   []PromotionDetail `json:"promotion_details,omitempty" xml:"promotion_details>promotion_detail,omitempty"`
    // 主订单包含的各个子订单信息
    SubOrders   []SubOrderInfo `json:"sub_orders,omitempty" xml:"sub_orders>sub_order_info,omitempty"`
    // 该笔订单是否押金合并支付订单（即该主订单是否已包含押金订单金额）
    OrderWithDepo   bool `json:"order_with_depo,omitempty" xml:"order_with_depo,omitempty"`
    // 订单邮费，需要邮寄的实体商品才有
    PostFee   int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
    // 评价对象
    ScoreDetails   []ScoreDetail `json:"score_details,omitempty" xml:"score_details>score_detail,omitempty"`
    // 取消原因
    CancelReason   string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
    // 评价xxxx
    RateContent   string `json:"rate_content,omitempty" xml:"rate_content,omitempty"`
}
