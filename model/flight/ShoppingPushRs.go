package flight

// ShoppingPushRs 
type ShoppingPushRs struct {

    // success
    Success   bool `json:"success,omitempty"`

    // errCode
    ErrCode   string `json:"err_code,omitempty"`

    // errMsg
    ErrMsg   string `json:"err_msg,omitempty"`

    // errRoutingMsg
    ErrRoutingMsgList   []Json `json:"err_routing_msg_list,omitempty"`

}
