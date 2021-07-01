package wdk

// PaymentItemDo 
type PaymentItemDo struct {
    // 商品sku码
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // 商品补差金额，单位为分
    PaymentFee   int64 `json:"payment_fee,omitempty" xml:"payment_fee,omitempty"`
}
