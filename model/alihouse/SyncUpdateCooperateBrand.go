package alihouse

import (
	"sync"
)

// SyncUpdateCooperateBrand 结构体
type SyncUpdateCooperateBrand struct {
	// 外部小区id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 外部房源id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部品牌id
	OuterBrandId string `json:"outer_brand_id,omitempty" xml:"outer_brand_id,omitempty"`
}

var poolSyncUpdateCooperateBrand = sync.Pool{
	New: func() any {
		return new(SyncUpdateCooperateBrand)
	},
}

// GetSyncUpdateCooperateBrand() 从对象池中获取SyncUpdateCooperateBrand
func GetSyncUpdateCooperateBrand() *SyncUpdateCooperateBrand {
	return poolSyncUpdateCooperateBrand.Get().(*SyncUpdateCooperateBrand)
}

// ReleaseSyncUpdateCooperateBrand 释放SyncUpdateCooperateBrand
func ReleaseSyncUpdateCooperateBrand(v *SyncUpdateCooperateBrand) {
	v.CommunityOuterId = ""
	v.OuterId = ""
	v.OuterBrandId = ""
	poolSyncUpdateCooperateBrand.Put(v)
}
