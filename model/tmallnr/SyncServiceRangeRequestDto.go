package tmallnr

import (
	"sync"
)

// SyncServiceRangeRequestDto 结构体
type SyncServiceRangeRequestDto struct {
	// 围栏信息
	Points []Point `json:"points,omitempty" xml:"points>point,omitempty"`
	// 扩展字段
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
	// 门店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 类型 -1 为删除,其余为添加，如果记录存在则做更新
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolSyncServiceRangeRequestDto = sync.Pool{
	New: func() any {
		return new(SyncServiceRangeRequestDto)
	},
}

// GetSyncServiceRangeRequestDto() 从对象池中获取SyncServiceRangeRequestDto
func GetSyncServiceRangeRequestDto() *SyncServiceRangeRequestDto {
	return poolSyncServiceRangeRequestDto.Get().(*SyncServiceRangeRequestDto)
}

// ReleaseSyncServiceRangeRequestDto 释放SyncServiceRangeRequestDto
func ReleaseSyncServiceRangeRequestDto(v *SyncServiceRangeRequestDto) {
	v.Points = v.Points[:0]
	v.Properties = ""
	v.StoreId = ""
	v.Type = 0
	poolSyncServiceRangeRequestDto.Put(v)
}
