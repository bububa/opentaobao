package tmallcms

// LLBApiResult 
/* model for simplify = false
type LLBApiResult struct {

    // 错误信息
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 结果对象
    
    Model  *struct {
        SpreadLinkDo  *SpreadLinkDo `json:"spread_link_do,omitempty"`
    } `json:"model,omitempty"`
    

    // 错误码
    
    ErrorCode   int64 `json:"error_code,omitempty"`
    

    // 成功结果
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// LLBApiResult 
type LLBApiResult struct {

    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty"`

    // 结果对象
    Model   *SpreadLinkDo `json:"model,omitempty"`

    // 错误码
    ErrorCode   int64 `json:"error_code,omitempty"`

    // 成功结果
    Success   bool `json:"success,omitempty"`

}
