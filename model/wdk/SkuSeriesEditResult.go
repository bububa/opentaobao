package wdk

// SkuSeriesEditResult 
type SkuSeriesEditResult struct {

    // 成功的商品编码集
    FailedSkuCodes   []String `json:"failed_sku_codes,omitempty"`

    // 失败的商品编码集
    SuccessedSkuCodes   []String `json:"successed_sku_codes,omitempty"`

    // 系列编码
    SeriesId   int64 `json:"series_id,omitempty"`

}
