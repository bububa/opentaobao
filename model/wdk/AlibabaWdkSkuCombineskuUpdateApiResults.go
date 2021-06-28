package wdk

// AlibabaWdkSkuCombineskuUpdateApiResults 
/* model for simplify = false
type AlibabaWdkSkuCombineskuUpdateApiResults struct {

    // 接口调用是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 接口调用异常编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 接口调用异常信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 商品列表
    
    Models  struct {
        AlibabaWdkSkuCombineskuUpdateApiResult  []AlibabaWdkSkuCombineskuUpdateApiResult `json:"alibaba_wdk_sku_combinesku_update_api_result,omitempty"`
    } `json:"models,omitempty"`
    

}
*/

// AlibabaWdkSkuCombineskuUpdateApiResults 
type AlibabaWdkSkuCombineskuUpdateApiResults struct {

    // 接口调用是否成功
    Success   bool `json:"success,omitempty"`

    // 接口调用异常编码
    ErrCode   string `json:"err_code,omitempty"`

    // 接口调用异常信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 商品列表
    Models   []AlibabaWdkSkuCombineskuUpdateApiResult `json:"models,omitempty"`

}
