package aedata

// Order 
type Order struct {

    // 佣金率
    
    CommissionRate   string `json:"commission_rate,omitempty" xml:"commission_rate,omitempty"`
    

    // 订单创建时间
    
    CreatedTime   string `json:"created_time,omitempty" xml:"created_time,omitempty"`
    

    // 自定义参数(JSON格式）
    
    CustomerParameters   string `json:"customer_parameters,omitempty" xml:"customer_parameters,omitempty"`
    

    // 订单状态:Payment Completed,Buyer Confirmed Receipt
    
    OrderStatus   string `json:"order_status,omitempty" xml:"order_status,omitempty"`
    

    // 订单完成时的预计佣金
    
    EstimatedFinishedCommission   int64 `json:"estimated_finished_commission,omitempty" xml:"estimated_finished_commission,omitempty"`
    

    // 订单支付时的预计佣金
    
    EstimatedPaidCommission   int64 `json:"estimated_paid_commission,omitempty" xml:"estimated_paid_commission,omitempty"`
    

    // 订单完成金额
    
    FinishedAmount   int64 `json:"finished_amount,omitempty" xml:"finished_amount,omitempty"`
    

    // 订单完成时间
    
    FinishedTime   string `json:"finished_time,omitempty" xml:"finished_time,omitempty"`
    

    // 是否新买家
    
    IsNewBuyer   string `json:"is_new_buyer,omitempty" xml:"is_new_buyer,omitempty"`
    

    // 下单商品数量
    
    ProductCount   int64 `json:"product_count,omitempty" xml:"product_count,omitempty"`
    

    // 商品DetailUrl
    
    ProductDetailUrl   string `json:"product_detail_url,omitempty" xml:"product_detail_url,omitempty"`
    

    // 商品ID
    
    ProductId   int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
    

    // 商品主图Url
    
    ProductMainImageUrl   string `json:"product_main_image_url,omitempty" xml:"product_main_image_url,omitempty"`
    

    // 商品标题
    
    ProductTitle   string `json:"product_title,omitempty" xml:"product_title,omitempty"`
    

    // 新买家奖励佣金
    
    NewBuyerBonusCommission   int64 `json:"new_buyer_bonus_commission,omitempty" xml:"new_buyer_bonus_commission,omitempty"`
    

    // 订单ID
    
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    

    // 子订单ID:已废弃，请使用sub_order_id
    
    OrderNumber   int64 `json:"order_number,omitempty" xml:"order_number,omitempty"`
    

    // 订单支付完成金额
    
    PaidAmount   int64 `json:"paid_amount,omitempty" xml:"paid_amount,omitempty"`
    

    // 订单支付完成时间
    
    PaidTime   string `json:"paid_time,omitempty" xml:"paid_time,omitempty"`
    

    // 父订单ID:已废弃，请使用order_id
    
    ParentOrderNumber   int64 `json:"parent_order_number,omitempty" xml:"parent_order_number,omitempty"`
    

    // 推广者结算币种
    
    SettledCurrency   string `json:"settled_currency,omitempty" xml:"settled_currency,omitempty"`
    

    // 订单物流国家
    
    ShipToCountry   string `json:"ship_to_country,omitempty" xml:"ship_to_country,omitempty"`
    

    // 子订单ID
    
    SubOrderId   int64 `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
    

    // trackId
    
    TrackingId   string `json:"tracking_id,omitempty" xml:"tracking_id,omitempty"`
    

    // 是否爆品订单:Y 或者N
    
    IsHotProduct   string `json:"is_hot_product,omitempty" xml:"is_hot_product,omitempty"`
    

    // 所属类目ID
    
    CategoryId   int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
    

    // FullRefund：全额退款，Antispam：风控拦截
    
    EffectDetailStatus   string `json:"effect_detail_status,omitempty" xml:"effect_detail_status,omitempty"`
    

    // 激励订单完成时的预计佣金
    
    EstimatedIncentiveFinishedCommission   int64 `json:"estimated_incentive_finished_commission,omitempty" xml:"estimated_incentive_finished_commission,omitempty"`
    

    // 激励订单支付时的预计佣金
    
    EstimatedIncentivePaidCommission   int64 `json:"estimated_incentive_paid_commission,omitempty" xml:"estimated_incentive_paid_commission,omitempty"`
    

    // 激励佣金率
    
    IncentiveCommissionRate   string `json:"incentive_commission_rate,omitempty" xml:"incentive_commission_rate,omitempty"`
    

    // 是否为联盟商品：N/Y
    
    IsAffiliateProduct   string `json:"is_affiliate_product,omitempty" xml:"is_affiliate_product,omitempty"`
    

    // 自定义参数(JSON格式）
    
    CustomParameters   string `json:"custom_parameters,omitempty" xml:"custom_parameters,omitempty"`
    

}
