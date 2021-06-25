package wdk

// SeriesSkuRequest 
type SeriesSkuRequest struct {

    // 系列编码
    SeriesId   int64 `json:"series_id,omitempty"`

    // 默认商品编码
    DefaultSkuCode   string `json:"default_sku_code,omitempty"`

    // 需要移除默认商品
    RemoveDefaultSku   bool `json:"remove_default_sku,omitempty"`

    // 商品编码集合
    SkuCodes   []String `json:"sku_codes,omitempty"`

}
