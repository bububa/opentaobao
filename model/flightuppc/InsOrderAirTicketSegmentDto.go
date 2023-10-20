package flightuppc

import (
	"sync"
)

// InsOrderAirTicketSegmentDto 结构体
type InsOrderAirTicketSegmentDto struct {
	// 票价
	TicketPrice string `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	// 航司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 到达城市
	ArrCity string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// 出发城市
	DepCity string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// 子保单号
	PolicyNo string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// 航班号
	FlightNo string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 起飞时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 到达时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

var poolInsOrderAirTicketSegmentDto = sync.Pool{
	New: func() any {
		return new(InsOrderAirTicketSegmentDto)
	},
}

// GetInsOrderAirTicketSegmentDto() 从对象池中获取InsOrderAirTicketSegmentDto
func GetInsOrderAirTicketSegmentDto() *InsOrderAirTicketSegmentDto {
	return poolInsOrderAirTicketSegmentDto.Get().(*InsOrderAirTicketSegmentDto)
}

// ReleaseInsOrderAirTicketSegmentDto 释放InsOrderAirTicketSegmentDto
func ReleaseInsOrderAirTicketSegmentDto(v *InsOrderAirTicketSegmentDto) {
	v.TicketPrice = ""
	v.CompanyName = ""
	v.ArrCity = ""
	v.DepCity = ""
	v.PolicyNo = ""
	v.FlightNo = ""
	v.TicketNo = ""
	v.StartTime = ""
	v.EndTime = ""
	poolInsOrderAirTicketSegmentDto.Put(v)
}
