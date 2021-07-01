package mos

// PosOrderDto 
type PosOrderDto struct {
    // POS交易流水号
    OutTradeNo   string `json:"out_trade_no,omitempty" xml:"out_trade_no,omitempty"`
    // 退款流水号（退款时必须有）
    OutRefundNo   string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
    // OC交易流水号
    TradeNo   string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
    // 会员卡号
    MemberCardNo   string `json:"member_card_no,omitempty" xml:"member_card_no,omitempty"`
    // 会员手机号
    MemberMobilePhone   string `json:"member_mobile_phone,omitempty" xml:"member_mobile_phone,omitempty"`
    // 门店号
    MallNo   string `json:"mall_no,omitempty" xml:"mall_no,omitempty"`
    // 订单金额，单位分
    TotalAmount   int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
    // 配送方式：1 : 门店自提，2 ：物流到家
    LogisticsWay   int64 `json:"logistics_way,omitempty" xml:"logistics_way,omitempty"`
    // 交易类别：0 : 收银  1 ：退款
    OrderType   int64 `json:"order_type,omitempty" xml:"order_type,omitempty"`
    // 操作员
    Operator   string `json:"operator,omitempty" xml:"operator,omitempty"`
    // 付款/退款 时间
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    // 订单商品列表
    SaleItems   []PosSaleItemDto `json:"sale_items,omitempty" xml:"sale_items>pos_sale_item_dto,omitempty"`
    // 支付列表
    Payments   []PosPaymentDto `json:"payments,omitempty" xml:"payments>pos_payment_dto,omitempty"`
    // 订单来源
    SaleChannle   int64 `json:"sale_channle,omitempty" xml:"sale_channle,omitempty"`
    // 收银机号
    TerminalNo   string `json:"terminal_no,omitempty" xml:"terminal_no,omitempty"`
    // 优惠分摊列表
    SplitPromotions   []PosSplitPromotionDto `json:"split_promotions,omitempty" xml:"split_promotions>pos_split_promotion_dto,omitempty"`
    // 扩展信息。
    ExtendParams   string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}
