package train

// ChangeTicketInfo 结构体
type ChangeTicketInfo struct {
	// 改签费
	ChangeFee int64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	// 扩展参数
	ExtendParam string `json:"extend_param,omitempty" xml:"extend_param,omitempty"`
	// 座位信息
	ChooseSeat string `json:"choose_seat,omitempty" xml:"choose_seat,omitempty"`
	// 子订单
	SubBizOrderId int64 `json:"sub_biz_order_id,omitempty" xml:"sub_biz_order_id,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 真实坐席
	RealSeat int64 `json:"real_seat,omitempty" xml:"real_seat,omitempty"`
	// 高改低手续费
	HandingFee int64 `json:"handing_fee,omitempty" xml:"handing_fee,omitempty"`
	// 是否支持在线退改签
	CanChange bool `json:"can_change,omitempty" xml:"can_change,omitempty"`
}
