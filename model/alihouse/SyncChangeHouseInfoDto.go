package alihouse

import (
	"sync"
)

// SyncChangeHouseInfoDto 结构体
type SyncChangeHouseInfoDto struct {
	// 老小区外部id
	OldCommunityOuterId string `json:"old_community_outer_id,omitempty" xml:"old_community_outer_id,omitempty"`
	// 小区外部id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// ecode编码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 1
	EntrustOuterId string `json:"entrust_outer_id,omitempty" xml:"entrust_outer_id,omitempty"`
	// 业务类型，业务类型，1-新房，2-二手房，3-租房，4房屋
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 租房业务模式 10-集中式模式 11-分散式模式
	HouseModel int64 `json:"house_model,omitempty" xml:"house_model,omitempty"`
	// 0-正常房源 1-临时房源（交易专用）
	HouseType int64 `json:"house_type,omitempty" xml:"house_type,omitempty"`
	// 是否异步，0同步，1异步
	IsAsync int64 `json:"is_async,omitempty" xml:"is_async,omitempty"`
	// 1
	IsChangeCommunity int64 `json:"is_change_community,omitempty" xml:"is_change_community,omitempty"`
}

var poolSyncChangeHouseInfoDto = sync.Pool{
	New: func() any {
		return new(SyncChangeHouseInfoDto)
	},
}

// GetSyncChangeHouseInfoDto() 从对象池中获取SyncChangeHouseInfoDto
func GetSyncChangeHouseInfoDto() *SyncChangeHouseInfoDto {
	return poolSyncChangeHouseInfoDto.Get().(*SyncChangeHouseInfoDto)
}

// ReleaseSyncChangeHouseInfoDto 释放SyncChangeHouseInfoDto
func ReleaseSyncChangeHouseInfoDto(v *SyncChangeHouseInfoDto) {
	v.OldCommunityOuterId = ""
	v.CommunityOuterId = ""
	v.OuterId = ""
	v.ECode = ""
	v.EntrustOuterId = ""
	v.BusinessType = 0
	v.HouseModel = 0
	v.HouseType = 0
	v.IsAsync = 0
	v.IsChangeCommunity = 0
	poolSyncChangeHouseInfoDto.Put(v)
}
