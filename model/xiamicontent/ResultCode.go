package xiamicontent

// ResultCode 
type ResultCode struct {
    // result message
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // result code
    Code   int64 `json:"code,omitempty" xml:"code,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
