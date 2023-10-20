package btrip

import (
	"sync"
)

// OpenInvoiceRuleRq 结构体
type OpenInvoiceRuleRq struct {
	// all_employe为true时非必传，否则必传
	Entities []Entity `json:"entities,omitempty" xml:"entities>entity,omitempty"`
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 第三方发票id
	ThirdPartId string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 是否适用所有员工
	AllEmploye bool `json:"all_employe,omitempty" xml:"all_employe,omitempty"`
}

var poolOpenInvoiceRuleRq = sync.Pool{
	New: func() any {
		return new(OpenInvoiceRuleRq)
	},
}

// GetOpenInvoiceRuleRq() 从对象池中获取OpenInvoiceRuleRq
func GetOpenInvoiceRuleRq() *OpenInvoiceRuleRq {
	return poolOpenInvoiceRuleRq.Get().(*OpenInvoiceRuleRq)
}

// ReleaseOpenInvoiceRuleRq 释放OpenInvoiceRuleRq
func ReleaseOpenInvoiceRuleRq(v *OpenInvoiceRuleRq) {
	v.Entities = v.Entities[:0]
	v.CorpId = ""
	v.ThirdPartId = ""
	v.Version = 0
	v.AllEmploye = false
	poolOpenInvoiceRuleRq.Put(v)
}
