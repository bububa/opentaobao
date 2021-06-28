package baichuan

// ErrorCode 
/* model for simplify = false
type ErrorCode struct {

    // 详细错误信息
    
    Message   string `json:"message,omitempty"`
    

    // 错误码
    
    Code   string `json:"code,omitempty"`
    

}
*/

// ErrorCode 
type ErrorCode struct {

    // 详细错误信息
    Message   string `json:"message,omitempty"`

    // 错误码
    Code   string `json:"code,omitempty"`

}
