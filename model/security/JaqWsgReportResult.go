package security

// JaqWsgReportResult 
/* model for simplify = false
type JaqWsgReportResult struct {

    // 生成的token
    
    Token   string `json:"token,omitempty"`
    

    // 错误码
    
    ErrorCode   int64 `json:"error_code,omitempty"`
    

    // 错误描述
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // 安全验证前向化下发参数结构体
    
    JaqDispatchParam  *struct {
        JaqDispatchParam  *JaqDispatchParam `json:"jaq_dispatch_param,omitempty"`
    } `json:"jaq_dispatch_param,omitempty"`
    

}
*/

// JaqWsgReportResult 
type JaqWsgReportResult struct {

    // 生成的token
    Token   string `json:"token,omitempty"`

    // 错误码
    ErrorCode   int64 `json:"error_code,omitempty"`

    // 错误描述
    ErrorMsg   string `json:"error_msg,omitempty"`

    // 安全验证前向化下发参数结构体
    JaqDispatchParam   *JaqDispatchParam `json:"jaq_dispatch_param,omitempty"`

}
