package ieagency

import (
	"sync"
)

// RefundItineraryVo 结构体
type RefundItineraryVo struct {
	// 航班列表
	RefundFlightSegmentVos []RefundFlightSegmentVo `json:"refund_flight_segment_vos,omitempty" xml:"refund_flight_segment_vos>refund_flight_segment_vo,omitempty"`
	// 到达机场
	ArrAirportCode string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// 出发机场
	DepAirportCode string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// 出发时间
	DepDate string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// 行程序号
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
}

var poolRefundItineraryVo = sync.Pool{
	New: func() any {
		return new(RefundItineraryVo)
	},
}

// GetRefundItineraryVo() 从对象池中获取RefundItineraryVo
func GetRefundItineraryVo() *RefundItineraryVo {
	return poolRefundItineraryVo.Get().(*RefundItineraryVo)
}

// ReleaseRefundItineraryVo 释放RefundItineraryVo
func ReleaseRefundItineraryVo(v *RefundItineraryVo) {
	v.RefundFlightSegmentVos = v.RefundFlightSegmentVos[:0]
	v.ArrAirportCode = ""
	v.DepAirportCode = ""
	v.DepDate = ""
	v.Index = 0
	poolRefundItineraryVo.Put(v)
}
