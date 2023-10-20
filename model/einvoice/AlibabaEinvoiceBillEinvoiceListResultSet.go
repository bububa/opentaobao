package einvoice

import (
	"sync"
)

// AlibabaEinvoiceBillEinvoiceListResultSet 结构体
type AlibabaEinvoiceBillEinvoiceListResultSet struct {
	// 返回结果具体信息
	ResultList []ResultList `json:"result_list,omitempty" xml:"result_list>result_list,omitempty"`
	// 错误信息
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolAlibabaEinvoiceBillEinvoiceListResultSet = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceBillEinvoiceListResultSet)
	},
}

// GetAlibabaEinvoiceBillEinvoiceListResultSet() 从对象池中获取AlibabaEinvoiceBillEinvoiceListResultSet
func GetAlibabaEinvoiceBillEinvoiceListResultSet() *AlibabaEinvoiceBillEinvoiceListResultSet {
	return poolAlibabaEinvoiceBillEinvoiceListResultSet.Get().(*AlibabaEinvoiceBillEinvoiceListResultSet)
}

// ReleaseAlibabaEinvoiceBillEinvoiceListResultSet 释放AlibabaEinvoiceBillEinvoiceListResultSet
func ReleaseAlibabaEinvoiceBillEinvoiceListResultSet(v *AlibabaEinvoiceBillEinvoiceListResultSet) {
	v.ResultList = v.ResultList[:0]
	v.ErrorMessage = ""
	v.RetCode = ""
	v.TotalCount = 0
	poolAlibabaEinvoiceBillEinvoiceListResultSet.Put(v)
}
