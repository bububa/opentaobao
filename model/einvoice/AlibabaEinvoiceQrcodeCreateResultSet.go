package einvoice

import (
	"sync"
)

// AlibabaEinvoiceQrcodeCreateResultSet 结构体
type AlibabaEinvoiceQrcodeCreateResultSet struct {
	// errorMessage
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
}

var poolAlibabaEinvoiceQrcodeCreateResultSet = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceQrcodeCreateResultSet)
	},
}

// GetAlibabaEinvoiceQrcodeCreateResultSet() 从对象池中获取AlibabaEinvoiceQrcodeCreateResultSet
func GetAlibabaEinvoiceQrcodeCreateResultSet() *AlibabaEinvoiceQrcodeCreateResultSet {
	return poolAlibabaEinvoiceQrcodeCreateResultSet.Get().(*AlibabaEinvoiceQrcodeCreateResultSet)
}

// ReleaseAlibabaEinvoiceQrcodeCreateResultSet 释放AlibabaEinvoiceQrcodeCreateResultSet
func ReleaseAlibabaEinvoiceQrcodeCreateResultSet(v *AlibabaEinvoiceQrcodeCreateResultSet) {
	v.ErrorMessage = ""
	v.Result = ""
	v.ErrorCode = ""
	poolAlibabaEinvoiceQrcodeCreateResultSet.Put(v)
}
