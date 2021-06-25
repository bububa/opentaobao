package bus

// B2BRefundOrderRp 
type B2BRefundOrderRp struct {

    // results
    Results   []Json `json:"results,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // success1
    Success1   bool `json:"success1,omitempty"`

}
