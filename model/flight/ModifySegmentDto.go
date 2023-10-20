package flight

import (
	"sync"
)

// ModifySegmentDto 结构体
type ModifySegmentDto struct {
	// 舱等: F:头等舱, C:商务舱, Y:经济舱, S:超级经济舱, P:超值经济舱, M:标准经济舱, W:超级经济舱
	CabinClass string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 起飞时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 起飞城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 舱位
	Cabin string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// 到达机场
	ArrAirport string `json:"arr_airport,omitempty" xml:"arr_airport,omitempty"`
	// 起飞机场
	DepAirport string `json:"dep_airport,omitempty" xml:"dep_airport,omitempty"`
	// 到达时间
	ArrTime string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// 该航段对应的票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 航段序号
	SegmentIndex int64 `json:"segment_index,omitempty" xml:"segment_index,omitempty"`
	// 航程序号
	OdIndex int64 `json:"od_index,omitempty" xml:"od_index,omitempty"`
}

var poolModifySegmentDto = sync.Pool{
	New: func() any {
		return new(ModifySegmentDto)
	},
}

// GetModifySegmentDto() 从对象池中获取ModifySegmentDto
func GetModifySegmentDto() *ModifySegmentDto {
	return poolModifySegmentDto.Get().(*ModifySegmentDto)
}

// ReleaseModifySegmentDto 释放ModifySegmentDto
func ReleaseModifySegmentDto(v *ModifySegmentDto) {
	v.CabinClass = ""
	v.FlightNo = ""
	v.DepTime = ""
	v.ArrCity = ""
	v.DepCity = ""
	v.Cabin = ""
	v.ArrAirport = ""
	v.DepAirport = ""
	v.ArrTime = ""
	v.TicketNo = ""
	v.SegmentIndex = 0
	v.OdIndex = 0
	poolModifySegmentDto.Put(v)
}
