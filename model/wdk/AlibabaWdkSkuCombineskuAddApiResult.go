package wdk

// AlibabaWdkSkuCombineskuAddApiResult 
/* model for simplify = false
type AlibabaWdkSkuCombineskuAddApiResult struct {

    // 单个商品是否新建成功
    
    Success   bool `json:"success,omitempty"`
    

    // 单个商品新建异常编码（异常才有值）
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 单个商品新建异常信息（异常才有值）
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 商品code值
    
    Model   string `json:"model,omitempty"`
    

}
*/

// AlibabaWdkSkuCombineskuAddApiResult 
type AlibabaWdkSkuCombineskuAddApiResult struct {

    // 单个商品是否新建成功
    Success   bool `json:"success,omitempty"`

    // 单个商品新建异常编码（异常才有值）
    ErrCode   string `json:"err_code,omitempty"`

    // 单个商品新建异常信息（异常才有值）
    ErrMsg   string `json:"err_msg,omitempty"`

    // 商品code值
    Model   string `json:"model,omitempty"`

}
