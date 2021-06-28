package trade

// ErrorCode 
/* model for simplify = false
type ErrorCode struct {

    // 错误展示信息
    
    DisplayMessage   string `json:"display_message,omitempty"`
    

    // 错误码key
    
    Key   string `json:"key,omitempty"`
    

    // 错误详情
    
    LogMessage   string `json:"log_message,omitempty"`
    

}
*/

// ErrorCode 
type ErrorCode struct {

    // 错误展示信息
    DisplayMessage   string `json:"display_message,omitempty"`

    // 错误码key
    Key   string `json:"key,omitempty"`

    // 错误详情
    LogMessage   string `json:"log_message,omitempty"`

}
