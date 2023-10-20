package btrip

import (
	"sync"
)

// BtripUserSyncRq 结构体
type BtripUserSyncRq struct {
	// 人员列表，最大长度5000
	UserList []UserSyncRq `json:"user_list,omitempty" xml:"user_list>user_sync_rq,omitempty"`
	// 第三方企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
}

var poolBtripUserSyncRq = sync.Pool{
	New: func() any {
		return new(BtripUserSyncRq)
	},
}

// GetBtripUserSyncRq() 从对象池中获取BtripUserSyncRq
func GetBtripUserSyncRq() *BtripUserSyncRq {
	return poolBtripUserSyncRq.Get().(*BtripUserSyncRq)
}

// ReleaseBtripUserSyncRq 释放BtripUserSyncRq
func ReleaseBtripUserSyncRq(v *BtripUserSyncRq) {
	v.UserList = v.UserList[:0]
	v.CorpId = ""
	poolBtripUserSyncRq.Put(v)
}
