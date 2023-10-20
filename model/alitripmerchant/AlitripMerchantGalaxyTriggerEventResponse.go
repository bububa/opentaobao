package alitripmerchant

// AlitripmerchantgalaxytriggereventResponse 结构体
type AlitripmerchantgalaxytriggereventResponse struct {
	// 状态码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 状态信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 消息体
	Content *EventCoupon `json:"content,omitempty" xml:"content,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
