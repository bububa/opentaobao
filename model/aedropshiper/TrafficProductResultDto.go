package aedropshiper

import (
	"sync"
)

// TrafficProductResultDto 结构体
type TrafficProductResultDto struct {
	// products
	Products []Integer `json:"products,omitempty" xml:"products>integer,omitempty"`
	// total record count
	TotalRecordCount int64 `json:"total_record_count,omitempty" xml:"total_record_count,omitempty"`
	// current record count
	CurrentRecordCount int64 `json:"current_record_count,omitempty" xml:"current_record_count,omitempty"`
	// total page number
	TotalPageNo int64 `json:"total_page_no,omitempty" xml:"total_page_no,omitempty"`
	// count page number
	CurrentPageNo int64 `json:"current_page_no,omitempty" xml:"current_page_no,omitempty"`
	// is finished
	IsFinished bool `json:"is_finished,omitempty" xml:"is_finished,omitempty"`
}

var poolTrafficProductResultDto = sync.Pool{
	New: func() any {
		return new(TrafficProductResultDto)
	},
}

// GetTrafficProductResultDto() 从对象池中获取TrafficProductResultDto
func GetTrafficProductResultDto() *TrafficProductResultDto {
	return poolTrafficProductResultDto.Get().(*TrafficProductResultDto)
}

// ReleaseTrafficProductResultDto 释放TrafficProductResultDto
func ReleaseTrafficProductResultDto(v *TrafficProductResultDto) {
	v.Products = v.Products[:0]
	v.TotalRecordCount = 0
	v.CurrentRecordCount = 0
	v.TotalPageNo = 0
	v.CurrentPageNo = 0
	v.IsFinished = false
	poolTrafficProductResultDto.Put(v)
}
