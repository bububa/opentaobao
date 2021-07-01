package lstlogistics2

// SendDummyOrderParam 结构体
type SendDummyOrderParam struct {
	// 发货时间
	SendTime string `json:"send_time,omitempty" xml:"send_time,omitempty"`
	// 备注
	Remarks string `json:"remarks,omitempty" xml:"remarks,omitempty"`
	// 发货主订单列表
	MainOrderParamList []MainOrderParam `json:"main_order_param_list,omitempty" xml:"main_order_param_list>main_order_param,omitempty"`
}
