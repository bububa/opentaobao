package btrip

// BtriphomeResultSupport 结构体
type BtriphomeResultSupport struct {
	// 审批单详情
	Apply *OpenApplyRs `json:"apply,omitempty" xml:"apply,omitempty"`
	// 成功标识
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// module
	ApplyList []OpenApplyRs `json:"apply_list,omitempty" xml:"apply_list>open_apply_rs,omitempty"`
	// 机票订单列表
	FlightOrderList []OpenFlightOrderRs `json:"flight_order_list,omitempty" xml:"flight_order_list>open_flight_order_rs,omitempty"`
	// 数据结果
	HotelOrderList []OpenHotelOrderRs `json:"hotel_order_list,omitempty" xml:"hotel_order_list>open_hotel_order_rs,omitempty"`
	// 火车票订单列表
	TrainOrderList []OpenTrainOrderRs `json:"train_order_list,omitempty" xml:"train_order_list>open_train_order_rs,omitempty"`
}
