package xhotelonlineorder

// StatementOrder 
type StatementOrder struct {

    // 淘宝订单ID
    
    Tid   int64 `json:"tid,omitempty" xml:"tid,omitempty"`
    

    // 外部订单ID
    
    OutId   string `json:"out_id,omitempty" xml:"out_id,omitempty"`
    

    // 淘宝佣金
    
    TaobaoCommission   string `json:"taobao_commission,omitempty" xml:"taobao_commission,omitempty"`
    

    // 交易状态（暂无值）
    
    TradeStatus   string `json:"trade_status,omitempty" xml:"trade_status,omitempty"`
    

    // 支付宝交易号
    
    AlipayTradeNo   string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
    

    // 总房费（分）
    
    Payment   string `json:"payment,omitempty" xml:"payment,omitempty"`
    

    // 扣除佣金后， 1.	集团入账=买家实际支付(房费+杂费-卖家优惠)； 2.	门店实际入账=买
    
    CommissionTotal   string `json:"commission_total,omitempty" xml:"commission_total,omitempty"`
    

    // 集团与门店佣金
    
    HotelCommission   string `json:"hotel_commission,omitempty" xml:"hotel_commission,omitempty"`
    

    // 卖家优惠（分）
    
    SellerPromotion   string `json:"seller_promotion,omitempty" xml:"seller_promotion,omitempty"`
    

    // 卖家优惠明细（暂无值）
    
    PromotionDetail   string `json:"promotion_detail,omitempty" xml:"promotion_detail,omitempty"`
    

    // 杂费总额
    
    OtherFee   string `json:"other_fee,omitempty" xml:"other_fee,omitempty"`
    

    // 结算日期
    
    SettleDate   string `json:"settle_date,omitempty" xml:"settle_date,omitempty"`
    

    // 入住天数
    
    RoomSumNights   int64 `json:"room_sum_nights,omitempty" xml:"room_sum_nights,omitempty"`
    

    // 分账状态 (0, "未分账"),(1, "分账成功"),(3, "分账失败"), (2, "无需分账"),(4,"预处理成功");
    
    SettleStatus   string `json:"settle_status,omitempty" xml:"settle_status,omitempty"`
    

    // 入住时间
    
    CheckIn   string `json:"check_in,omitempty" xml:"check_in,omitempty"`
    

    // 离店时间
    
    CheckOut   string `json:"check_out,omitempty" xml:"check_out,omitempty"`
    

    // 门店收款账户(1：支付宝；2：银行卡)
    
    AccountType   string `json:"account_type,omitempty" xml:"account_type,omitempty"`
    

    // 支付类型(1:预付；5:面付;7：在线预约;601:线上信用住；602：扫码信用住；603：身份证线下信用住；604：官网信用住)
    
    PayType   string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
    

    // 实际的房型名称
    
    RoomTypeName   string `json:"room_type_name,omitempty" xml:"room_type_name,omitempty"`
    

    // 交易手续费
    
    TransactionFee   string `json:"transaction_fee,omitempty" xml:"transaction_fee,omitempty"`
    

    // 酒店名称
    
    HotelName   string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
    

    // 税和费，单位分
    
    TaxAndFee   string `json:"tax_and_fee,omitempty" xml:"tax_and_fee,omitempty"`
    

}
