package wdk

// OctopusOpenResult 
/* model for simplify = false
type OctopusOpenResult struct {

    // 活动ID
    
    Data   int64 `json:"data,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 操作是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 错误描述
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 部分失败的商品编码列表
    
    FailedSkuCodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"failed_sku_codes,omitempty"`
    

}
*/

// OctopusOpenResult 
type OctopusOpenResult struct {

    // 活动ID
    Data   int64 `json:"data,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 操作是否成功
    Success   bool `json:"success,omitempty"`

    // 错误描述
    ErrorMessage   string `json:"error_message,omitempty"`

    // 部分失败的商品编码列表
    FailedSkuCodes   []string `json:"failed_sku_codes,omitempty"`

}
