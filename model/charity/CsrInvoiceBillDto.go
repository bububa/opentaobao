package charity

import (
	"sync"
)

// CsrInvoiceBillDto 结构体
type CsrInvoiceBillDto struct {
	// 账单明细文件地址，行数据中携带制表符
	BillFile string `json:"bill_file,omitempty" xml:"bill_file,omitempty"`
	// 账期
	BillCycle string `json:"bill_cycle,omitempty" xml:"bill_cycle,omitempty"`
	// 账单明细文件地址,行数据中不带制表符
	BillDetailStandardFile string `json:"bill_detail_standard_file,omitempty" xml:"bill_detail_standard_file,omitempty"`
}

var poolCsrInvoiceBillDto = sync.Pool{
	New: func() any {
		return new(CsrInvoiceBillDto)
	},
}

// GetCsrInvoiceBillDto() 从对象池中获取CsrInvoiceBillDto
func GetCsrInvoiceBillDto() *CsrInvoiceBillDto {
	return poolCsrInvoiceBillDto.Get().(*CsrInvoiceBillDto)
}

// ReleaseCsrInvoiceBillDto 释放CsrInvoiceBillDto
func ReleaseCsrInvoiceBillDto(v *CsrInvoiceBillDto) {
	v.BillFile = ""
	v.BillCycle = ""
	v.BillDetailStandardFile = ""
	poolCsrInvoiceBillDto.Put(v)
}
