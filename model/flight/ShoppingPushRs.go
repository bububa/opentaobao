package flight

// ShoppingPushRs 
type ShoppingPushRs struct {

    // success
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // errCode
    
    ErrCode   string `json:"err_code,omitempty" xml:"err_code,omitempty"`
    

    // errMsg
    
    ErrMsg   string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
    

    // errRoutingMsg
    
    ErrRoutingMsgList   []string `json:"err_routing_msg_list,omitempty" xml:"err_routing_msg_list>string,omitempty"`
    

}
