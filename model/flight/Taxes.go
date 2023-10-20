package flight

import (
	"sync"
)

// Taxes 结构体
type Taxes struct {
	// 税项二字码
	TaxCode string `json:"tax_code,omitempty" xml:"tax_code,omitempty"`
	// 税值
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolTaxes = sync.Pool{
	New: func() any {
		return new(Taxes)
	},
}

// GetTaxes() 从对象池中获取Taxes
func GetTaxes() *Taxes {
	return poolTaxes.Get().(*Taxes)
}

// ReleaseTaxes 释放Taxes
func ReleaseTaxes(v *Taxes) {
	v.TaxCode = ""
	v.Amount = 0
	poolTaxes.Put(v)
}
