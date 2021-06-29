package exchange

// RefundBaseResponse 
type RefundBaseResponse struct {

    // message
    
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
    

    // msgCode
    
    MsgCode   string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // exchange
    
    Exchange   *Exchange `json:"exchange,omitempty" xml:"exchange,omitempty"`
    

}
