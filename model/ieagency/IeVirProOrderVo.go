package ieagency

// IeVirProOrderVo 结构体
type IeVirProOrderVo struct {
	// 辅营订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 辅营购买信息
	PassengerAuxVos []IePassengerAuxVo `json:"passenger_aux_vos,omitempty" xml:"passenger_aux_vos>ie_passenger_aux_vo,omitempty"`
}
