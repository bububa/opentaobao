package wdk

// AlibabaWdkSkuCombineskuQueryApiResult 
/* model for simplify = false
type AlibabaWdkSkuCombineskuQueryApiResult struct {

    // 单个商品是否查询成功
    
    Success   bool `json:"success,omitempty"`
    

    // 单个商品查询异常编码（异常才有值）
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 单个商品查询异常信息（异常才有值）
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 商品信息
    
    Model  *struct {
        SkuDo  *SkuDo `json:"sku_do,omitempty"`
    } `json:"model,omitempty"`
    

}
*/

// AlibabaWdkSkuCombineskuQueryApiResult 
type AlibabaWdkSkuCombineskuQueryApiResult struct {

    // 单个商品是否查询成功
    Success   bool `json:"success,omitempty"`

    // 单个商品查询异常编码（异常才有值）
    ErrCode   string `json:"err_code,omitempty"`

    // 单个商品查询异常信息（异常才有值）
    ErrMsg   string `json:"err_msg,omitempty"`

    // 商品信息
    Model   *SkuDo `json:"model,omitempty"`

}
