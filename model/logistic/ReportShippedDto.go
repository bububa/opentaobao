package logistic

import (
	"sync"
)

// ReportShippedDto 结构体
type ReportShippedDto struct {
	// shipment order id
	LogisticsOrderId string `json:"logistics_order_id,omitempty" xml:"logistics_order_id,omitempty"`
	// type of report dispatch
	ShippedType string `json:"shipped_type,omitempty" xml:"shipped_type,omitempty"`
}

var poolReportShippedDto = sync.Pool{
	New: func() any {
		return new(ReportShippedDto)
	},
}

// GetReportShippedDto() 从对象池中获取ReportShippedDto
func GetReportShippedDto() *ReportShippedDto {
	return poolReportShippedDto.Get().(*ReportShippedDto)
}

// ReleaseReportShippedDto 释放ReportShippedDto
func ReleaseReportShippedDto(v *ReportShippedDto) {
	v.LogisticsOrderId = ""
	v.ShippedType = ""
	poolReportShippedDto.Put(v)
}
