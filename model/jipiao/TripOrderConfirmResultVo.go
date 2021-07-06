package jipiao

// TripOrderConfirmResultVo 结构体
type TripOrderConfirmResultVo struct {
	// 返回结果，此接口无值
	Results []string `json:"results,omitempty" xml:"results>string,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 订单是否成功
	IsOrderSuccess bool `json:"is_order_success,omitempty" xml:"is_order_success,omitempty"`
}
