package wlb

// ResultDo 
type ResultDo struct {
    // 网络延时
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // 01
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // 成功、失败
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
