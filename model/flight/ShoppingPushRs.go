package flight

// ShoppingPushRs 
/* model for simplify = false
type ShoppingPushRs struct {

    // success
    
    Success   bool `json:"success,omitempty"`
    

    // errCode
    
    ErrCode   string `json:"err_code,omitempty"`
    

    // errMsg
    
    ErrMsg   string `json:"err_msg,omitempty"`
    

    // errRoutingMsg
    
    ErrRoutingMsgList  struct {
        Json  []string `json:"string,omitempty"`
    } `json:"err_routing_msg_list,omitempty"`
    

}
*/

// ShoppingPushRs 
type ShoppingPushRs struct {

    // success
    Success   bool `json:"success,omitempty"`

    // errCode
    ErrCode   string `json:"err_code,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

    // errRoutingMsg
    ErrRoutingMsgList   []string `json:"err_routing_msg_list,omitempty"`

}
