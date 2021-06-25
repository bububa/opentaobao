package wdk

// ItemConfirmInfo 
type ItemConfirmInfo struct {

    // 确认数量(为正数或零)
    ConfirmQuantity   string `json:"confirm_quantity,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

}
