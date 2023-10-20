package btrip

import (
	"sync"
)

// OpenInvoiceRq 结构体
type OpenInvoiceRq struct {
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 发票抬头关键词
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 第三方用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 2：商旅开放平台
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOpenInvoiceRq = sync.Pool{
	New: func() any {
		return new(OpenInvoiceRq)
	},
}

// GetOpenInvoiceRq() 从对象池中获取OpenInvoiceRq
func GetOpenInvoiceRq() *OpenInvoiceRq {
	return poolOpenInvoiceRq.Get().(*OpenInvoiceRq)
}

// ReleaseOpenInvoiceRq 释放OpenInvoiceRq
func ReleaseOpenInvoiceRq(v *OpenInvoiceRq) {
	v.CorpId = ""
	v.Title = ""
	v.UserId = ""
	v.Version = 0
	poolOpenInvoiceRq.Put(v)
}
