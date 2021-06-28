package wdk

// AlibabaWdkSeriesSkuRemoveApiResult 
/* model for simplify = false
type AlibabaWdkSeriesSkuRemoveApiResult struct {

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 系列品移除商品成功结果
    
    Model  *struct {
        SkuSeriesEditResult  *SkuSeriesEditResult `json:"sku_series_edit_result,omitempty"`
    } `json:"model,omitempty"`
    

    // 错误详情
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaWdkSeriesSkuRemoveApiResult 
type AlibabaWdkSeriesSkuRemoveApiResult struct {

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 系列品移除商品成功结果
    Model   *SkuSeriesEditResult `json:"model,omitempty"`

    // 错误详情
    ErrMsg   string `json:"err_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
