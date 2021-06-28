package wdk

// AlibabaWdkSkuQueryApiResult 
/* model for simplify = false
type AlibabaWdkSkuQueryApiResult struct {

    // 单条查询结果
    
    Model  *struct {
        SkuDo  *SkuDo `json:"sku_do,omitempty"`
    } `json:"model,omitempty"`
    

    // 请求参数不能为空
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // 单条错误信息
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 单条是否成功
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// AlibabaWdkSkuQueryApiResult 
type AlibabaWdkSkuQueryApiResult struct {

    // 单条查询结果
    Model   *SkuDo `json:"model,omitempty"`

    // 请求参数不能为空
    ErrCode   string `json:"err_code,omitempty"`

    // 单条错误信息
    ErrMsg   string `json:"err_msg,omitempty"`

    // 单条是否成功
    Success   bool `json:"success,omitempty"`

}
