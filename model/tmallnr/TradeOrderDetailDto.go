package tmallnr

// TradeOrderDetailDto 
type TradeOrderDetailDto struct {
    // 子订单列表
    OrderDTOs   []NrOrderDto `json:"order_d_t_os,omitempty" xml:"order_d_t_os>nr_order_dto,omitempty"`
    // 门店编码
    StoreCode   string `json:"store_code,omitempty" xml:"store_code,omitempty"`
    // 预约结束时间
    AppointEndTime   string `json:"appoint_end_time,omitempty" xml:"appoint_end_time,omitempty"`
    // 预约开始时间
    AppointStartTime   string `json:"appoint_start_time,omitempty" xml:"appoint_start_time,omitempty"`
    // 收货详细地址_商家没有入塔不提供
    ReceiverAddress   string `json:"receiver_address,omitempty" xml:"receiver_address,omitempty"`
    // 镇_商家没有入塔不提供
    ReceiverTown   string `json:"receiver_town,omitempty" xml:"receiver_town,omitempty"`
    // 区_商家没有入塔不提供
    ReceiverDistrict   string `json:"receiver_district,omitempty" xml:"receiver_district,omitempty"`
    // 城市_商家没有入塔不提供
    ReceiverCity   string `json:"receiver_city,omitempty" xml:"receiver_city,omitempty"`
    // 省_商家没有入塔不提供
    ReceiverProvince   string `json:"receiver_province,omitempty" xml:"receiver_province,omitempty"`
    // 电话_商家没有入塔不提供
    ReceiverPhone   string `json:"receiver_phone,omitempty" xml:"receiver_phone,omitempty"`
    // 手机_商家没有入塔不提供
    ReceiverMobile   string `json:"receiver_mobile,omitempty" xml:"receiver_mobile,omitempty"`
    // 收货人_商家没有入塔不提供
    ReceiverName   string `json:"receiver_name,omitempty" xml:"receiver_name,omitempty"`
    // 支付时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 交易状态可选值:TRADE_NO_CREATE_PAY(没有创建支付宝交易)，WAIT_BUYER_PAY(等待买家付款)，SELLER_CONSIGNED_PART(卖家部分发货)，WAIT_SELLER_SEND_GOODS(等待卖家发货,即:买家已付款)，WAIT_BUYER_CONFIRM_GOODS(等待买家确认收货,即:卖家已发货)，TRADE_BUYER_SIGNED(买家已签收,货到付款专用)，TRADE_FINISHED(交易成功)，TRADE_CLOSED(付款以后用户退款成功，交易自动关闭)，TRADE_CLOSED_BY_TAOBAO(付款以前，卖家或买家主动关闭交易)，PAY_PENDING(国际信用卡支付付款确认中) * WAIT_PRE_AUTH_CONFIRM(0元购合约中)，PAID_FORBID_CONSIGN(拼团中订单或者发货强管控的订单，已付款但禁止发货)
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 主订单
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // createTime
    CreateTime   string `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 买家昵称
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    // 卖家昵称
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    // 门店的外部编号
    OutIdStoreCode   string `json:"out_id_store_code,omitempty" xml:"out_id_store_code,omitempty"`
    // 邮费（单位分）
    PostFee   string `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
    // 买家留言
    BuyerMemo   string `json:"buyer_memo,omitempty" xml:"buyer_memo,omitempty"`
}
