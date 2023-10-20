package logistic

import (
	"sync"
)

// ReportShippedRequestDto 结构体
type ReportShippedRequestDto struct {
	// report shipped list
	ReportShippedList []ReportShippedDto `json:"report_shipped_list,omitempty" xml:"report_shipped_list>report_shipped_dto,omitempty"`
}

var poolReportShippedRequestDto = sync.Pool{
	New: func() any {
		return new(ReportShippedRequestDto)
	},
}

// GetReportShippedRequestDto() 从对象池中获取ReportShippedRequestDto
func GetReportShippedRequestDto() *ReportShippedRequestDto {
	return poolReportShippedRequestDto.Get().(*ReportShippedRequestDto)
}

// ReleaseReportShippedRequestDto 释放ReportShippedRequestDto
func ReleaseReportShippedRequestDto(v *ReportShippedRequestDto) {
	v.ReportShippedList = v.ReportShippedList[:0]
	poolReportShippedRequestDto.Put(v)
}
