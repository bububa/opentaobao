package util

// QimenResult 
/* model for simplify = false
type QimenResult struct {

    // 错误信息
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

}
*/

// QimenResult 
type QimenResult struct {

    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

}
