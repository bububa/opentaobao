package wdk

// AlibabaWdkSkuCombineskuQueryApiResults 
/* model for simplify = false
type AlibabaWdkSkuCombineskuQueryApiResults struct {

    // 接口调用是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 接口调用异常编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 接口调用异常信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 商品列表
    
    Models  struct {
        AlibabaWdkSkuCombineskuQueryApiResult  []AlibabaWdkSkuCombineskuQueryApiResult `json:"alibaba_wdk_sku_combinesku_query_api_result,omitempty"`
    } `json:"models,omitempty"`
    

}
*/

// AlibabaWdkSkuCombineskuQueryApiResults 
type AlibabaWdkSkuCombineskuQueryApiResults struct {

    // 接口调用是否成功
    Success   bool `json:"success,omitempty"`

    // 接口调用异常编码
    ErrCode   string `json:"err_code,omitempty"`

    // 接口调用异常信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 商品列表
    Models   []AlibabaWdkSkuCombineskuQueryApiResult `json:"models,omitempty"`

}
