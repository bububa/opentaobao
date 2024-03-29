package einvoice

import (
	"sync"
)

// EinvoiceAmountCheckResult 结构体
type EinvoiceAmountCheckResult struct {
	// 税号
	PayeeRegisterNo string `json:"payee_register_no,omitempty" xml:"payee_register_no,omitempty"`
	// 开票日期
	InvoiceDate string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
	// 开票含税总金额
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 开票不含税总金额
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 开票总税额
	TotalTax string `json:"total_tax,omitempty" xml:"total_tax,omitempty"`
	// 发票类型，蓝票=blue，红票=red
	InvoiceType string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 开票量
	InvoiceCount int64 `json:"invoice_count,omitempty" xml:"invoice_count,omitempty"`
}

var poolEinvoiceAmountCheckResult = sync.Pool{
	New: func() any {
		return new(EinvoiceAmountCheckResult)
	},
}

// GetEinvoiceAmountCheckResult() 从对象池中获取EinvoiceAmountCheckResult
func GetEinvoiceAmountCheckResult() *EinvoiceAmountCheckResult {
	return poolEinvoiceAmountCheckResult.Get().(*EinvoiceAmountCheckResult)
}

// ReleaseEinvoiceAmountCheckResult 释放EinvoiceAmountCheckResult
func ReleaseEinvoiceAmountCheckResult(v *EinvoiceAmountCheckResult) {
	v.PayeeRegisterNo = ""
	v.InvoiceDate = ""
	v.TotalAmount = ""
	v.TotalPrice = ""
	v.TotalTax = ""
	v.InvoiceType = ""
	v.InvoiceCount = 0
	poolEinvoiceAmountCheckResult.Put(v)
}
