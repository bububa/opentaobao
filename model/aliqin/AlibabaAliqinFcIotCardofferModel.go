package aliqin

// AlibabaAliqinFcIotCardofferModel 
type AlibabaAliqinFcIotCardofferModel struct {

    // 失效时间
    ExpireTime   string `json:"expire_time,omitempty"`

    // 生效时间
    EffectiveTime   string `json:"effective_time,omitempty"`

    // 订购时间
    OrderTime   string `json:"order_time,omitempty"`

    // 商品名称
    OfferName   string `json:"offer_name,omitempty"`

    // 商品ID
    OfferId   string `json:"offer_id,omitempty"`

}
