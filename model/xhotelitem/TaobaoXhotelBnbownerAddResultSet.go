package xhotelitem

// TaobaoXhotelBnbownerAddResultSet 
type TaobaoXhotelBnbownerAddResultSet struct {
    // 是否出错
    Error   bool `json:"error,omitempty" xml:"error,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // firstResult
    FirstResult   *AddOwnerParam `json:"first_result,omitempty" xml:"first_result,omitempty"`
    // 是否成功标记
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
