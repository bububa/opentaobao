package alsc

import (
	"sync"
)

// TaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult 结构体
type TaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult struct {
	// 鹰眼id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回结果数据
	Data *TopInvoiceInsuranceNoDto `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult = sync.Pool{
	New: func() any {
		return new(TaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult)
	},
}

// GetTaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult() 从对象池中获取TaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult
func GetTaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult() *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult {
	return poolTaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult.Get().(*TaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult)
}

// ReleaseTaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult 释放TaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult
func ReleaseTaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult(v *TaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult) {
	v.TraceId = ""
	v.Data = nil
	v.Error = nil
	v.Success = false
	poolTaobaoServindustryFinanceInsuranceInvoiceInsurancenosResult.Put(v)
}
