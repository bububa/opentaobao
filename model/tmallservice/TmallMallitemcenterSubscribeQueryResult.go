package tmallservice

// TmallMallitemcenterSubscribeQueryResult 
/* model for simplify = false
type TmallMallitemcenterSubscribeQueryResult struct {

    // 返回结果model
    
    ResultData  *struct {
        ResultData  *ResultData `json:"result_data,omitempty"`
    } `json:"result_data,omitempty"`
    

    // message
    
    Message   string `json:"message,omitempty"`
    

    // 错误类型
    
    ErrorType   int64 `json:"error_type,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // 是否正常
    
    Success   bool `json:"success,omitempty"`
    

}
*/

// TmallMallitemcenterSubscribeQueryResult 
type TmallMallitemcenterSubscribeQueryResult struct {

    // 返回结果model
    ResultData   *ResultData `json:"result_data,omitempty"`

    // message
    Message   string `json:"message,omitempty"`

    // 错误类型
    ErrorType   int64 `json:"error_type,omitempty"`

    // 错误码
    ErrorCode   string `json:"error_code,omitempty"`

    // 是否正常
    Success   bool `json:"success,omitempty"`

}
