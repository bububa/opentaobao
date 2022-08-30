package flight

// ShoppingPushRs 结构体
type ShoppingPushRs struct {
	// errRoutingMsg
	ErrRoutingMsgList []string `json:"err_routing_msg_list,omitempty" xml:"err_routing_msg_list>string,omitempty"`
	// errMsg
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
