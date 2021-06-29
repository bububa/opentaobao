package hotelalliance

// HmsTopResultSet 
type HmsTopResultSet struct {

    // 返回module
    
    Result   *AllianceInfoResult `json:"result,omitempty" xml:"result,omitempty"`
    

    // 操作是否成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 错误code
    
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

}
