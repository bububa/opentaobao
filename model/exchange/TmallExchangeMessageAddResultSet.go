package exchange

// TmallExchangeMessageAddResultSet 
type TmallExchangeMessageAddResultSet struct {

    // 留言信息
    
    Results   []RefundMessage `json:"results,omitempty" xml:"results,omitempty"`
    

    // 异常信息
    
    Exception   string `json:"exception,omitempty" xml:"exception,omitempty"`
    

    // 错误码
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

}
