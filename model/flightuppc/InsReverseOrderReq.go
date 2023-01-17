package flightuppc

// InsReverseOrderReq 结构体
type InsReverseOrderReq struct {
	// 保险信息列表
	Insureds []InsPersonAndAirSegmentDto `json:"insureds,omitempty" xml:"insureds>ins_person_and_air_segment_dto,omitempty"`
	// 保险订单号
	TpOrderId int64 `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
}
