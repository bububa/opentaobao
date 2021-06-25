package alicom

// OpentradCreateOrderRequestDTO 
type OpentradCreateOrderRequestDTO struct {

    // 活动ID
    ActivityId   int64 `json:"activity_id,omitempty"`

    // 交易请求流水号
    TransferId   string `json:"transfer_id,omitempty"`

    // 手机号
    Phone   string `json:"phone,omitempty"`

    // 价格(分)
    Price   string `json:"price,omitempty"`

    // 来源
    Source   string `json:"source,omitempty"`

    // 卖家Nick
    SellerNick   string `json:"seller_nick,omitempty"`

    // 支付宝Id
    AlipayId   int64 `json:"alipay_id,omitempty"`

    // 淘宝Nick
    TaobaoNick   string `json:"taobao_nick,omitempty"`

    // 淘宝Token
    TaobaoToken   string `json:"taobao_token,omitempty"`

    // 礼包ID
    GiftId   int64 `json:"gift_id,omitempty"`

    // 产品名称
    ProductName   string `json:"product_name,omitempty"`

    // 扩展属性
    Ext   string `json:"ext,omitempty"`

    // 产品ID
    ProductId   string `json:"product_id,omitempty"`

}
