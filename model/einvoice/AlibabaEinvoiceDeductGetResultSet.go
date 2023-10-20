package einvoice

import (
	"sync"
)

// AlibabaEinvoiceDeductGetResultSet 结构体
type AlibabaEinvoiceDeductGetResultSet struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 业务日期
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 税号
	PackageRegisterNo string `json:"package_register_no,omitempty" xml:"package_register_no,omitempty"`
	// result
	Result *AlibabaEinvoiceDeductGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
	// totalCount
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 实际扣减
	Deduct int64 `json:"deduct,omitempty" xml:"deduct,omitempty"`
	// 应扣减
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolAlibabaEinvoiceDeductGetResultSet = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceDeductGetResultSet)
	},
}

// GetAlibabaEinvoiceDeductGetResultSet() 从对象池中获取AlibabaEinvoiceDeductGetResultSet
func GetAlibabaEinvoiceDeductGetResultSet() *AlibabaEinvoiceDeductGetResultSet {
	return poolAlibabaEinvoiceDeductGetResultSet.Get().(*AlibabaEinvoiceDeductGetResultSet)
}

// ReleaseAlibabaEinvoiceDeductGetResultSet 释放AlibabaEinvoiceDeductGetResultSet
func ReleaseAlibabaEinvoiceDeductGetResultSet(v *AlibabaEinvoiceDeductGetResultSet) {
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.BizDate = ""
	v.PackageRegisterNo = ""
	v.Result = nil
	v.TotalCount = 0
	v.Deduct = 0
	v.Amount = 0
	poolAlibabaEinvoiceDeductGetResultSet.Put(v)
}
