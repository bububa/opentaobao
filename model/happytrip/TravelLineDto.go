package happytrip

import (
	"sync"
)

// TravelLineDto 结构体
type TravelLineDto struct {
	// 开始时间
	BeginDate string `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
	// 结束时间
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 出发城市
	FromCity string `json:"from_city,omitempty" xml:"from_city,omitempty"`
	// 行程类型，ONE_WAY：单程，ROUND_TRIP：往返，MULTI_WAY：多程
	ItineraryType string `json:"itinerary_type,omitempty" xml:"itinerary_type,omitempty"`
	// 目的地城市
	ToCity string `json:"to_city,omitempty" xml:"to_city,omitempty"`
	// 交通类型，PLANE：飞机，TRAIN：火车，AUTOMOTIVE：汽车，STEAMSHIP：轮船，OTHER：其他
	TransportType string `json:"transport_type,omitempty" xml:"transport_type,omitempty"`
	// 行程说明
	TravelPurpose string `json:"travel_purpose,omitempty" xml:"travel_purpose,omitempty"`
}

var poolTravelLineDto = sync.Pool{
	New: func() any {
		return new(TravelLineDto)
	},
}

// GetTravelLineDto() 从对象池中获取TravelLineDto
func GetTravelLineDto() *TravelLineDto {
	return poolTravelLineDto.Get().(*TravelLineDto)
}

// ReleaseTravelLineDto 释放TravelLineDto
func ReleaseTravelLineDto(v *TravelLineDto) {
	v.BeginDate = ""
	v.EndDate = ""
	v.FromCity = ""
	v.ItineraryType = ""
	v.ToCity = ""
	v.TransportType = ""
	v.TravelPurpose = ""
	poolTravelLineDto.Put(v)
}
