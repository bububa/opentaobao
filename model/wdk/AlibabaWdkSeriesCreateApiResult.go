package wdk

// AlibabaWdkSeriesCreateApiResult 
/* model for simplify = false
type AlibabaWdkSeriesCreateApiResult struct {

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 商品系列创建结果
    
    Model  *struct {
        SkuSeriesCreateResult  *SkuSeriesCreateResult `json:"sku_series_create_result,omitempty"`
    } `json:"model,omitempty"`
    

    // 错误详情
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaWdkSeriesCreateApiResult 
type AlibabaWdkSeriesCreateApiResult struct {

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 商品系列创建结果
    Model   *SkuSeriesCreateResult `json:"model,omitempty"`

    // 错误详情
    ErrMsg   string `json:"err_msg,omitempty"`

    // 是否成功
    Success   bool `json:"success,omitempty"`

}
