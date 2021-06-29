package tmallnr

// NrOrderDTO 
type NrOrderDTO struct {
    // 创建时间
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 买家昵称
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    // 卖家昵称
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    // 退款状态退款状态。退款状态。可选值WAIT_SELLER_AGREE(买家已经申请退款，等待卖家同意)，WAIT_BUYER_RETURN_GOODS(卖家已经同意退款，等待买家退货)，WAIT_SELLER_CONFIRM_GOODS(买家已经退货，等待卖家确认收货)，SELLER_REFUSE_BUYER(卖家拒绝退款)，CLOSED(退款关闭)，SUCCESS(退款成功)
    RefundStatus   string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
    // 购买数量
    Num   int64 `json:"num,omitempty" xml:"num,omitempty"`
    // skuId
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // 商品ID
    ItemId   int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 商品标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 子订单号
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 修改价格修改的金额
    AdjustFee   int64 `json:"adjust_fee,omitempty" xml:"adjust_fee,omitempty"`
    // 店铺优惠的分摊
    DiscountFee   int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
    // 实际付款金额
    ActualPaidFee   int64 `json:"actual_paid_fee,omitempty" xml:"actual_paid_fee,omitempty"`
    // 商品价格
    AuctionPrice   int64 `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
    // outIdItemCode
    OutIdItemCode   string `json:"out_id_item_code,omitempty" xml:"out_id_item_code,omitempty"`
    // 配送计划的详情，仅做周期送业务需要
    NrZqsPlanRespDTO   *NrZqsPlanRespDTO `json:"nr_zqs_plan_resp_d_t_o,omitempty" xml:"nr_zqs_plan_resp_d_t_o,omitempty"`
    // 商家sku的外部编码
    OuterIdSku   string `json:"outer_id_sku,omitempty" xml:"outer_id_sku,omitempty"`
}
