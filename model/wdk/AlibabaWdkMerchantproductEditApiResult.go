package wdk

// AlibabaWdkMerchantproductEditApiResult 
/* model for simplify = false
type AlibabaWdkMerchantproductEditApiResult struct {

    // 请求是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误码
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 返回结果
    
    Model  *struct {
        MerchantProductResponse  *MerchantProductResponse `json:"merchant_product_response,omitempty"`
    } `json:"model,omitempty"`
    

}
*/

// AlibabaWdkMerchantproductEditApiResult 
type AlibabaWdkMerchantproductEditApiResult struct {

    // 请求是否成功
    Success   bool `json:"success,omitempty"`

    // 错误码
    ErrCode   string `json:"err_code,omitempty"`

    // 错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 返回结果
    Model   *MerchantProductResponse `json:"model,omitempty"`

}
