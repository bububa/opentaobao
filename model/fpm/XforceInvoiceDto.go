package fpm

import (
	"sync"
)

// XforceInvoiceDto 结构体
type XforceInvoiceDto struct {
	// 发票行
	DetailList []InvoiceDetails `json:"detail_list,omitempty" xml:"detail_list>invoice_details,omitempty"`
	// 发票头
	Head *InvoiceMainExt `json:"head,omitempty" xml:"head,omitempty"`
}

var poolXforceInvoiceDto = sync.Pool{
	New: func() any {
		return new(XforceInvoiceDto)
	},
}

// GetXforceInvoiceDto() 从对象池中获取XforceInvoiceDto
func GetXforceInvoiceDto() *XforceInvoiceDto {
	return poolXforceInvoiceDto.Get().(*XforceInvoiceDto)
}

// ReleaseXforceInvoiceDto 释放XforceInvoiceDto
func ReleaseXforceInvoiceDto(v *XforceInvoiceDto) {
	v.DetailList = v.DetailList[:0]
	v.Head = nil
	poolXforceInvoiceDto.Put(v)
}
