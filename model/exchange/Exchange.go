package exchange

// Exchange 
type Exchange struct {

    // 换货单状态
    
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    

    // 修改时间
    
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    

    // 换货单号ID
    
    DisputeId   string `json:"dispute_id,omitempty" xml:"dispute_id,omitempty"`
    

    // 正向交易单号ID
    
    BizOrderId   string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
    

    // 买家发货物流单号
    
    BuyerLogisticNo   string `json:"buyer_logistic_no,omitempty" xml:"buyer_logistic_no,omitempty"`
    

    // 支付宝单号ID
    
    AlipayNo   string `json:"alipay_no,omitempty" xml:"alipay_no,omitempty"`
    

    // 换货理由说明
    
    Desc   string `json:"desc,omitempty" xml:"desc,omitempty"`
    

    // 换货申请理由
    
    Reason   string `json:"reason,omitempty" xml:"reason,omitempty"`
    

    // 申请换货的状态：售中换货 or 售后换货
    
    RefundPhase   string `json:"refund_phase,omitempty" xml:"refund_phase,omitempty"`
    

    // 操作场景
    
    OperationContraint   string `json:"operation_contraint,omitempty" xml:"operation_contraint,omitempty"`
    

    // 商品名称
    
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    

    // 换货单创建时间
    
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    

    // 卖家昵称
    
    SellerNick   string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
    

    // 先行垫付状态
    
    AdvanceStatus   int64 `json:"advance_status,omitempty" xml:"advance_status,omitempty"`
    

    // 买家昵称
    
    BuyerNick   string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
    

    // 换货版本
    
    RefundVersion   int64 `json:"refund_version,omitempty" xml:"refund_version,omitempty"`
    

    // 换货数量
    
    Num   int64 `json:"num,omitempty" xml:"num,omitempty"`
    

    // 卖家发货快递单号
    
    SellerLogisticNo   string `json:"seller_logistic_no,omitempty" xml:"seller_logistic_no,omitempty"`
    

    // 价格
    
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    

    // 小二托管状态
    
    CsStatus   int64 `json:"cs_status,omitempty" xml:"cs_status,omitempty"`
    

    // 卖家换货地址
    
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    

    // 商品状态
    
    GoodStatus   string `json:"good_status,omitempty" xml:"good_status,omitempty"`
    

    // 换货商品的sku
    
    ExchangeSku   string `json:"exchange_sku,omitempty" xml:"exchange_sku,omitempty"`
    

    // buyerAddress
    
    BuyerAddress   string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
    

    // 超时时间
    
    TimeOut   string `json:"time_out,omitempty" xml:"time_out,omitempty"`
    

    // 买家发货物流公司名称
    
    BuyerLogisticName   string `json:"buyer_logistic_name,omitempty" xml:"buyer_logistic_name,omitempty"`
    

    // 卖家发货物流公司名称
    
    SellerLogisticName   string `json:"seller_logistic_name,omitempty" xml:"seller_logistic_name,omitempty"`
    

    // 所购买的商品sku
    
    BoughtSku   string `json:"bought_sku,omitempty" xml:"bought_sku,omitempty"`
    

    // 买家联系方式
    
    BuyerPhone   string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
    

    // buyerName
    
    BuyerName   string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
    

    // attributes
    
    Attributes   string `json:"attributes,omitempty" xml:"attributes,omitempty"`
    

    // 支付费用
    
    Payment   string `json:"payment,omitempty" xml:"payment,omitempty"`
    

}
