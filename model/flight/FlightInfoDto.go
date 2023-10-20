package flight

import (
	"sync"
)

// FlightInfoDto 结构体
type FlightInfoDto struct {
	// 航段信息
	SegmentInfos []SegmentInfoDto `json:"segment_infos,omitempty" xml:"segment_infos>segment_info_dto,omitempty"`
	// 航司
	Airway string `json:"airway,omitempty" xml:"airway,omitempty"`
	// 抵达机场
	ArrAirportCode string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// 起飞机场
	DepAirportCode string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// 第一段起飞时间
	DepDateTime string `json:"dep_date_time,omitempty" xml:"dep_date_time,omitempty"`
}

var poolFlightInfoDto = sync.Pool{
	New: func() any {
		return new(FlightInfoDto)
	},
}

// GetFlightInfoDto() 从对象池中获取FlightInfoDto
func GetFlightInfoDto() *FlightInfoDto {
	return poolFlightInfoDto.Get().(*FlightInfoDto)
}

// ReleaseFlightInfoDto 释放FlightInfoDto
func ReleaseFlightInfoDto(v *FlightInfoDto) {
	v.SegmentInfos = v.SegmentInfos[:0]
	v.Airway = ""
	v.ArrAirportCode = ""
	v.DepAirportCode = ""
	v.DepDateTime = ""
	poolFlightInfoDto.Put(v)
}
