package xhotelonlineorder

import (
	"sync"
)

// TaobaoXhotelCommoninvoiceListVtwoResultSet 结构体
type TaobaoXhotelCommoninvoiceListVtwoResultSet struct {
	// 常用发票信息
	Results []CommonInvoiceInfo `json:"results,omitempty" xml:"results>common_invoice_info,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 是否成功标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoXhotelCommoninvoiceListVtwoResultSet = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelCommoninvoiceListVtwoResultSet)
	},
}

// GetTaobaoXhotelCommoninvoiceListVtwoResultSet() 从对象池中获取TaobaoXhotelCommoninvoiceListVtwoResultSet
func GetTaobaoXhotelCommoninvoiceListVtwoResultSet() *TaobaoXhotelCommoninvoiceListVtwoResultSet {
	return poolTaobaoXhotelCommoninvoiceListVtwoResultSet.Get().(*TaobaoXhotelCommoninvoiceListVtwoResultSet)
}

// ReleaseTaobaoXhotelCommoninvoiceListVtwoResultSet 释放TaobaoXhotelCommoninvoiceListVtwoResultSet
func ReleaseTaobaoXhotelCommoninvoiceListVtwoResultSet(v *TaobaoXhotelCommoninvoiceListVtwoResultSet) {
	v.Results = v.Results[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.Success = false
	poolTaobaoXhotelCommoninvoiceListVtwoResultSet.Put(v)
}
