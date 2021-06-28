package wdk

// SkuSeriesEditResult 
/* model for simplify = false
type SkuSeriesEditResult struct {

    // 成功的商品编码集
    
    FailedSkuCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"failed_sku_codes,omitempty"`
    

    // 失败的商品编码集
    
    SuccessedSkuCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"successed_sku_codes,omitempty"`
    

    // 系列编码
    
    SeriesId   int64 `json:"series_id,omitempty"`
    

}
*/

// SkuSeriesEditResult 
type SkuSeriesEditResult struct {

    // 成功的商品编码集
    FailedSkuCodes   []string `json:"failed_sku_codes,omitempty"`

    // 失败的商品编码集
    SuccessedSkuCodes   []string `json:"successed_sku_codes,omitempty"`

    // 系列编码
    SeriesId   int64 `json:"series_id,omitempty"`

}
