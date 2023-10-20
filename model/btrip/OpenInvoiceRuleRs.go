package btrip

import (
	"sync"
)

// OpenInvoiceRuleRs 结构体
type OpenInvoiceRuleRs struct {
	// 新增数
	AddNum int64 `json:"add_num,omitempty" xml:"add_num,omitempty"`
	// 删除数
	RemoveNum int64 `json:"remove_num,omitempty" xml:"remove_num,omitempty"`
}

var poolOpenInvoiceRuleRs = sync.Pool{
	New: func() any {
		return new(OpenInvoiceRuleRs)
	},
}

// GetOpenInvoiceRuleRs() 从对象池中获取OpenInvoiceRuleRs
func GetOpenInvoiceRuleRs() *OpenInvoiceRuleRs {
	return poolOpenInvoiceRuleRs.Get().(*OpenInvoiceRuleRs)
}

// ReleaseOpenInvoiceRuleRs 释放OpenInvoiceRuleRs
func ReleaseOpenInvoiceRuleRs(v *OpenInvoiceRuleRs) {
	v.AddNum = 0
	v.RemoveNum = 0
	poolOpenInvoiceRuleRs.Put(v)
}
