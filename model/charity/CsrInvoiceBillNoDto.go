package charity

import (
	"sync"
)

// CsrInvoiceBillNoDto 结构体
type CsrInvoiceBillNoDto struct {
	// 账单编号
	BillId string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
	// 账单时间
	BillTime int64 `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
}

var poolCsrInvoiceBillNoDto = sync.Pool{
	New: func() any {
		return new(CsrInvoiceBillNoDto)
	},
}

// GetCsrInvoiceBillNoDto() 从对象池中获取CsrInvoiceBillNoDto
func GetCsrInvoiceBillNoDto() *CsrInvoiceBillNoDto {
	return poolCsrInvoiceBillNoDto.Get().(*CsrInvoiceBillNoDto)
}

// ReleaseCsrInvoiceBillNoDto 释放CsrInvoiceBillNoDto
func ReleaseCsrInvoiceBillNoDto(v *CsrInvoiceBillNoDto) {
	v.BillId = ""
	v.BillTime = 0
	poolCsrInvoiceBillNoDto.Put(v)
}
