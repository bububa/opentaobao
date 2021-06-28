package bus

// B2BRefundOrderRp 
/* model for simplify = false
type B2BRefundOrderRp struct {

    // results
    
    Results  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"results,omitempty"`
    

    // errorCode
    
    ErrorCode   string `json:"error_code,omitempty"`
    

    // errorMsg
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // success1
    
    Success1   bool `json:"success1,omitempty"`
    

}
*/

// B2BRefundOrderRp 
type B2BRefundOrderRp struct {

    // results
    Results   []string `json:"results,omitempty"`

    // errorCode
    ErrorCode   string `json:"error_code,omitempty"`

    // errorMsg
    ErrorMsg   string `json:"error_msg,omitempty"`

    // success
    Success   bool `json:"success,omitempty"`

    // success1
    Success1   bool `json:"success1,omitempty"`

}
