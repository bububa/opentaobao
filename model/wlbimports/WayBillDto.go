package wlbimports

import (
	"sync"
)

// WayBillDto 结构体
type WayBillDto struct {
	// 云打印数据
	CloudPrintData string `json:"cloud_print_data,omitempty" xml:"cloud_print_data,omitempty"`
	// 云打印pdf
	PdfWayBillUrl string `json:"pdf_way_bill_url,omitempty" xml:"pdf_way_bill_url,omitempty"`
}

var poolWayBillDto = sync.Pool{
	New: func() any {
		return new(WayBillDto)
	},
}

// GetWayBillDto() 从对象池中获取WayBillDto
func GetWayBillDto() *WayBillDto {
	return poolWayBillDto.Get().(*WayBillDto)
}

// ReleaseWayBillDto 释放WayBillDto
func ReleaseWayBillDto(v *WayBillDto) {
	v.CloudPrintData = ""
	v.PdfWayBillUrl = ""
	poolWayBillDto.Put(v)
}
