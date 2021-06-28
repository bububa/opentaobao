package trade

// TribeError 
/* model for simplify = false
type TribeError struct {

    // 错误可读性描述
    
    View   string `json:"view,omitempty"`
    

    // 错误码
    
    Code   string `json:"code,omitempty"`
    

    // 错误信息
    
    Message   string `json:"message,omitempty"`
    

}
*/

// TribeError 
type TribeError struct {

    // 错误可读性描述
    View   string `json:"view,omitempty"`

    // 错误码
    Code   string `json:"code,omitempty"`

    // 错误信息
    Message   string `json:"message,omitempty"`

}
