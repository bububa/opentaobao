package bus

// B2BRefundOrderRp 
type B2BRefundOrderRp struct {
    // results
    Results   []string `json:"results,omitempty" xml:"results>string,omitempty"`
    // errorCode
    ErrorCode   string `json:"error_code,omitempty" xml:"error_code,omitempty"`
    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    // success
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // success1
    Success1   bool `json:"success1,omitempty" xml:"success1,omitempty"`
}
