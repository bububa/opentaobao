package flightuppc

import (
	"sync"
)

// FlightChangeDto 结构体
type FlightChangeDto struct {
	// 航班号
	OldFlightNo string `json:"old_flight_no,omitempty" xml:"old_flight_no,omitempty"`
	// 原航班日期
	OldFlightDate string `json:"old_flight_date,omitempty" xml:"old_flight_date,omitempty"`
	// 出发机场三字码
	OldDepartCode string `json:"old_depart_code,omitempty" xml:"old_depart_code,omitempty"`
	// 到达机场三字码
	OldArriveCode string `json:"old_arrive_code,omitempty" xml:"old_arrive_code,omitempty"`
	// 出发时间
	OldDepartTime string `json:"old_depart_time,omitempty" xml:"old_depart_time,omitempty"`
	// 到达时间
	OldArriveTime string `json:"old_arrive_time,omitempty" xml:"old_arrive_time,omitempty"`
	// 新航班号
	NewFlightNo string `json:"new_flight_no,omitempty" xml:"new_flight_no,omitempty"`
	// 新航班日期
	NewFlightDate string `json:"new_flight_date,omitempty" xml:"new_flight_date,omitempty"`
	// 新出发机场三字码
	NewDepartCode string `json:"new_depart_code,omitempty" xml:"new_depart_code,omitempty"`
	// 新到达机场三字码
	NewArriveCode string `json:"new_arrive_code,omitempty" xml:"new_arrive_code,omitempty"`
	// 新出发时间
	NewDepartTime string `json:"new_depart_time,omitempty" xml:"new_depart_time,omitempty"`
	// 新到达时间
	NewArriveTime string `json:"new_arrive_time,omitempty" xml:"new_arrive_time,omitempty"`
	// 航变时间
	ChangeTime string `json:"change_time,omitempty" xml:"change_time,omitempty"`
	// 航变原因
	ChangeReason string `json:"change_reason,omitempty" xml:"change_reason,omitempty"`
	// 航变类型,1为取消，2为变更
	ChangeType int64 `json:"change_type,omitempty" xml:"change_type,omitempty"`
	// 航变子原因，如延误、恢复、航班保护，详情见新京杭航变接入文档
	SubType int64 `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
}

var poolFlightChangeDto = sync.Pool{
	New: func() any {
		return new(FlightChangeDto)
	},
}

// GetFlightChangeDto() 从对象池中获取FlightChangeDto
func GetFlightChangeDto() *FlightChangeDto {
	return poolFlightChangeDto.Get().(*FlightChangeDto)
}

// ReleaseFlightChangeDto 释放FlightChangeDto
func ReleaseFlightChangeDto(v *FlightChangeDto) {
	v.OldFlightNo = ""
	v.OldFlightDate = ""
	v.OldDepartCode = ""
	v.OldArriveCode = ""
	v.OldDepartTime = ""
	v.OldArriveTime = ""
	v.NewFlightNo = ""
	v.NewFlightDate = ""
	v.NewDepartCode = ""
	v.NewArriveCode = ""
	v.NewDepartTime = ""
	v.NewArriveTime = ""
	v.ChangeTime = ""
	v.ChangeReason = ""
	v.ChangeType = 0
	v.SubType = 0
	poolFlightChangeDto.Put(v)
}
