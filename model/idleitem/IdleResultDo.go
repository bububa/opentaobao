package idleitem

// IdleResultDo 
type IdleResultDo struct {
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 商品id
    Data   int64 `json:"data,omitempty" xml:"data,omitempty"`
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
