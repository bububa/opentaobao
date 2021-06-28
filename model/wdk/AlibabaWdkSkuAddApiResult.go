package wdk

// AlibabaWdkSkuAddApiResult 
/* model for simplify = false
type AlibabaWdkSkuAddApiResult struct {

    // 错误编码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 单个商品新增是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // skuCode商品编码
    
    Model   string `json:"model,omitempty"`
    

    // 聚合之后的产品id，商家需要保存该值
    
    ProductId   string `json:"product_id,omitempty"`
    

}
*/

// AlibabaWdkSkuAddApiResult 
type AlibabaWdkSkuAddApiResult struct {

    // 错误编码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 单个商品新增是否成功
    Success   bool `json:"success,omitempty"`

    // skuCode商品编码
    Model   string `json:"model,omitempty"`

    // 聚合之后的产品id，商家需要保存该值
    ProductId   string `json:"product_id,omitempty"`

}
