package btrip

import (
	"sync"
)

// BtripDepartSyncRq 结构体
type BtripDepartSyncRq struct {
	// 部门列表
	DepartList []DepartSyncRq `json:"depart_list,omitempty" xml:"depart_list>depart_sync_rq,omitempty"`
	// 第三方企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
}

var poolBtripDepartSyncRq = sync.Pool{
	New: func() any {
		return new(BtripDepartSyncRq)
	},
}

// GetBtripDepartSyncRq() 从对象池中获取BtripDepartSyncRq
func GetBtripDepartSyncRq() *BtripDepartSyncRq {
	return poolBtripDepartSyncRq.Get().(*BtripDepartSyncRq)
}

// ReleaseBtripDepartSyncRq 释放BtripDepartSyncRq
func ReleaseBtripDepartSyncRq(v *BtripDepartSyncRq) {
	v.DepartList = v.DepartList[:0]
	v.CorpId = ""
	poolBtripDepartSyncRq.Put(v)
}
