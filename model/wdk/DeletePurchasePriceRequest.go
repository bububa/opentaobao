package wdk

// DeletePurchasePriceRequest 
type DeletePurchasePriceRequest struct {

    // 请求幂等ID
    OutId   string `json:"out_id,omitempty"`

    // 商品skucode
    SkuCode   string `json:"sku_code,omitempty"`

    // 门店ID
    OuCode   string `json:"ou_code,omitempty"`

}
