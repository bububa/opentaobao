package tmallsc

// Resultdata 
type Resultdata struct {

    // 账单ID
    
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    

    // 服务单号
    
    ServiceOrderId   int64 `json:"service_order_id,omitempty" xml:"service_order_id,omitempty"`
    

    // 工单ID
    
    WorkcardId   int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
    

    // 工单多次作业时的批次号
    
    WorkcardSequence   int64 `json:"workcard_sequence,omitempty" xml:"workcard_sequence,omitempty"`
    

    // 卖家nick
    
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    

    // 交易主订单号
    
    ParentTradeOrderId   int64 `json:"parent_trade_order_id,omitempty" xml:"parent_trade_order_id,omitempty"`
    

    // 交易实物订单号
    
    MasterTradeOrderId   int64 `json:"master_trade_order_id,omitempty" xml:"master_trade_order_id,omitempty"`
    

    // 交易服务订单号
    
    ServiceTradeOrderId   int64 `json:"service_trade_order_id,omitempty" xml:"service_trade_order_id,omitempty"`
    

    // 支付宝交易订单号
    
    AlipayOrderId   string `json:"alipay_order_id,omitempty" xml:"alipay_order_id,omitempty"`
    

    // 门店id
    
    ServiceStoreId   int64 `json:"service_store_id,omitempty" xml:"service_store_id,omitempty"`
    

    // 门店Code
    
    ServiceStoreCode   string `json:"service_store_code,omitempty" xml:"service_store_code,omitempty"`
    

    // 门店名称
    
    ServiceStoreName   string `json:"service_store_name,omitempty" xml:"service_store_name,omitempty"`
    

    // 转帐金额，单位分
    
    TransferAmount   int64 `json:"transfer_amount,omitempty" xml:"transfer_amount,omitempty"`
    

    // 创建时间，单位毫秒
    
    CreateTime   int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
    

    // 打款时间，单位毫秒
    
    PayTime   int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
    

    // 打款备注
    
    Comment   string `json:"comment,omitempty" xml:"comment,omitempty"`
    

    // 货币类型 人民币(CNY)
    
    Currency   string `json:"currency,omitempty" xml:"currency,omitempty"`
    

    // 入款方支付宝账号
    
    InUserAlipayAccountId   string `json:"in_user_alipay_account_id,omitempty" xml:"in_user_alipay_account_id,omitempty"`
    

    // 状态；销帐：FINISH ;未销账:ONGOING
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 入款方nick
    
    InUserNick   string `json:"in_user_nick,omitempty" xml:"in_user_nick,omitempty"`
    

    // 入款方用户角色 BUYER:买家 SELLER:卖家, STORE：门店, TP：服务商, PLATFORM：平台
    
    InUserRole   string `json:"in_user_role,omitempty" xml:"in_user_role,omitempty"`
    

    // 出款方支付宝账号
    
    OutUserAlipayAccountId   string `json:"out_user_alipay_account_id,omitempty" xml:"out_user_alipay_account_id,omitempty"`
    

    // 出款方nick
    
    OutUserNick   string `json:"out_user_nick,omitempty" xml:"out_user_nick,omitempty"`
    

    // 入款方用户角色 BUYER:买家 SELLER:卖家, STORE：门店, TP：服务商, PLATFORM：平台
    
    OutUserRole   string `json:"out_user_role,omitempty" xml:"out_user_role,omitempty"`
    

    // 扩展信息；json格式
    
    Attributes   string `json:"attributes,omitempty" xml:"attributes,omitempty"`
    

    // 触发本次打款的动作类型。serviceOrder-tradeSuccess: 商家扣款到中间账户；serviceOrder-cancel: 退款;huijin-commision: 门店抽佣;huijin-store-transfer: 转账给门店
    
    Action   string `json:"action,omitempty" xml:"action,omitempty"`
    

}
