package mtopopen

// AlibabaInteractAopdataRegisterResult 
/* model for simplify = false
type AlibabaInteractAopdataRegisterResult struct {

    // xx
    
    Data   string `json:"data,omitempty"`
    

    // xxx失败
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // xxx失败
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // 接口调用成功
    
    Success   bool `json:"success,omitempty"`
    

    // 跟踪请求使用
    
    TraceId   string `json:"trace_id,omitempty"`
    

}
*/

// AlibabaInteractAopdataRegisterResult 
type AlibabaInteractAopdataRegisterResult struct {

    // xx
    Data   string `json:"data,omitempty"`

    // xxx失败
    ErrCode   string `json:"err_code,omitempty"`

    // xxx失败
    ErrMsg   string `json:"err_msg,omitempty"`

    // 接口调用成功
    Success   bool `json:"success,omitempty"`

    // 跟踪请求使用
    TraceId   string `json:"trace_id,omitempty"`

}
