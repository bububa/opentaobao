package btrip

import (
	"sync"
)

// OpenInvoiceDeleteRq 结构体
type OpenInvoiceDeleteRq struct {
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 第三方发票id
	ThirdPartId string `json:"third_part_id,omitempty" xml:"third_part_id,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOpenInvoiceDeleteRq = sync.Pool{
	New: func() any {
		return new(OpenInvoiceDeleteRq)
	},
}

// GetOpenInvoiceDeleteRq() 从对象池中获取OpenInvoiceDeleteRq
func GetOpenInvoiceDeleteRq() *OpenInvoiceDeleteRq {
	return poolOpenInvoiceDeleteRq.Get().(*OpenInvoiceDeleteRq)
}

// ReleaseOpenInvoiceDeleteRq 释放OpenInvoiceDeleteRq
func ReleaseOpenInvoiceDeleteRq(v *OpenInvoiceDeleteRq) {
	v.CorpId = ""
	v.ThirdPartId = ""
	v.Version = 0
	poolOpenInvoiceDeleteRq.Put(v)
}
