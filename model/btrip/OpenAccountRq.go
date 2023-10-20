package btrip

import (
	"sync"
)

// OpenAccountRq 结构体
type OpenAccountRq struct {
	// 对账单月份，不传取最新对账单
	BillMonth string `json:"bill_month,omitempty" xml:"bill_month,omitempty"`
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOpenAccountRq = sync.Pool{
	New: func() any {
		return new(OpenAccountRq)
	},
}

// GetOpenAccountRq() 从对象池中获取OpenAccountRq
func GetOpenAccountRq() *OpenAccountRq {
	return poolOpenAccountRq.Get().(*OpenAccountRq)
}

// ReleaseOpenAccountRq 释放OpenAccountRq
func ReleaseOpenAccountRq(v *OpenAccountRq) {
	v.BillMonth = ""
	v.CorpId = ""
	v.Version = 0
	poolOpenAccountRq.Put(v)
}
