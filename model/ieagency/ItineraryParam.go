package ieagency

import (
	"sync"
)

// ItineraryParam 结构体
type ItineraryParam struct {
	// 航段信息
	FlightSegmentParams []FlightSegmentParam `json:"flight_segment_params,omitempty" xml:"flight_segment_params>flight_segment_param,omitempty"`
	// 航程序号（从1开始）
	ItineraryRph int64 `json:"itinerary_rph,omitempty" xml:"itinerary_rph,omitempty"`
}

var poolItineraryParam = sync.Pool{
	New: func() any {
		return new(ItineraryParam)
	},
}

// GetItineraryParam() 从对象池中获取ItineraryParam
func GetItineraryParam() *ItineraryParam {
	return poolItineraryParam.Get().(*ItineraryParam)
}

// ReleaseItineraryParam 释放ItineraryParam
func ReleaseItineraryParam(v *ItineraryParam) {
	v.FlightSegmentParams = v.FlightSegmentParams[:0]
	v.ItineraryRph = 0
	poolItineraryParam.Put(v)
}
