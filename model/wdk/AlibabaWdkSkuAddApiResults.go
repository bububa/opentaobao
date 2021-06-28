package wdk

// AlibabaWdkSkuAddApiResults 
/* model for simplify = false
type AlibabaWdkSkuAddApiResults struct {

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 接口返回成功标志
    
    Success   bool `json:"success,omitempty"`
    

    // models
    
    Models  struct {
        AlibabaWdkSkuAddApiResult  []AlibabaWdkSkuAddApiResult `json:"alibaba_wdk_sku_add_api_result,omitempty"`
    } `json:"models,omitempty"`
    

}
*/

// AlibabaWdkSkuAddApiResults 
type AlibabaWdkSkuAddApiResults struct {

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 接口返回成功标志
    Success   bool `json:"success,omitempty"`

    // models
    Models   []AlibabaWdkSkuAddApiResult `json:"models,omitempty"`

}
