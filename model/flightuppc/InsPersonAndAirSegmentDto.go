package flightuppc

import (
	"sync"
)

// InsPersonAndAirSegmentDto 结构体
type InsPersonAndAirSegmentDto struct {
	// 航段信息
	InsOrderSegment *InsOrderAirTicketSegmentDto `json:"ins_order_segment,omitempty" xml:"ins_order_segment,omitempty"`
	// 被保人
	InsPerson *InsPersonDto `json:"ins_person,omitempty" xml:"ins_person,omitempty"`
}

var poolInsPersonAndAirSegmentDto = sync.Pool{
	New: func() any {
		return new(InsPersonAndAirSegmentDto)
	},
}

// GetInsPersonAndAirSegmentDto() 从对象池中获取InsPersonAndAirSegmentDto
func GetInsPersonAndAirSegmentDto() *InsPersonAndAirSegmentDto {
	return poolInsPersonAndAirSegmentDto.Get().(*InsPersonAndAirSegmentDto)
}

// ReleaseInsPersonAndAirSegmentDto 释放InsPersonAndAirSegmentDto
func ReleaseInsPersonAndAirSegmentDto(v *InsPersonAndAirSegmentDto) {
	v.InsOrderSegment = nil
	v.InsPerson = nil
	poolInsPersonAndAirSegmentDto.Put(v)
}
