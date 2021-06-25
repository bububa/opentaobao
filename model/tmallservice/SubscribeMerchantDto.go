package tmallservice

// SubscribeMerchantDto 
type SubscribeMerchantDto struct {

    // 技术电话
    TechPhone   string `json:"tech_phone,omitempty"`

    // 操作人电话
    OperatorPhone   string `json:"operator_phone,omitempty"`

    // 售后电话
    AfterSalePhone   string `json:"after_sale_phone,omitempty"`

    // 订购单主键id
    SubscribeOrderId   int64 `json:"subscribe_order_id,omitempty"`

    // spuid
    SpuId   int64 `json:"spu_id,omitempty"`

    // 服务cide
    ServiceCode   string `json:"service_code,omitempty"`

    // 订购者昵称
    SubscriberNick   string `json:"subscriber_nick,omitempty"`

    // 卖家类型
    SellerType   string `json:"seller_type,omitempty"`

}
