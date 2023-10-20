package alsc

import (
	"sync"
)

// TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult 结构体
type TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult struct {
	// 鹰眼id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 开票结果
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoServindustryFinanceInsuranceInvoiceFeedbackResult = sync.Pool{
	New: func() any {
		return new(TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult)
	},
}

// GetTaobaoServindustryFinanceInsuranceInvoiceFeedbackResult() 从对象池中获取TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult
func GetTaobaoServindustryFinanceInsuranceInvoiceFeedbackResult() *TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult {
	return poolTaobaoServindustryFinanceInsuranceInvoiceFeedbackResult.Get().(*TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult)
}

// ReleaseTaobaoServindustryFinanceInsuranceInvoiceFeedbackResult 释放TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult
func ReleaseTaobaoServindustryFinanceInsuranceInvoiceFeedbackResult(v *TaobaoServindustryFinanceInsuranceInvoiceFeedbackResult) {
	v.TraceId = ""
	v.Data = ""
	v.Error = nil
	v.Success = false
	poolTaobaoServindustryFinanceInsuranceInvoiceFeedbackResult.Put(v)
}
