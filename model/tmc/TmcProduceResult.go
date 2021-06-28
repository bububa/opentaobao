package tmc

// TmcProduceResult 
/* model for simplify = false
type TmcProduceResult struct {

    // 错误信息
    
    ErrorMessage   string `json:"error_message,omitempty"`
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

}
*/

// TmcProduceResult 
type TmcProduceResult struct {

    // 错误信息
    ErrorMessage   string `json:"error_message,omitempty"`

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

}
