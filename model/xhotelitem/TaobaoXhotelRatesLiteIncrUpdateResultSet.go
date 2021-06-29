package xhotelitem

// TaobaoXhotelRatesLiteIncrUpdateResultSet 
type TaobaoXhotelRatesLiteIncrUpdateResultSet struct {
    // 多个rate的更新结果
    FirstResult   string `json:"first_result,omitempty" xml:"first_result,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
