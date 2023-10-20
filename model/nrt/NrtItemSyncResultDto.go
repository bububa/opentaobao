package nrt

import (
	"sync"
)

// NrtItemSyncResultDto 结构体
type NrtItemSyncResultDto struct {
	// 摊位商品ID
	SItemId int64 `json:"s_item_id,omitempty" xml:"s_item_id,omitempty"`
	// 主商品ID
	MItemId int64 `json:"m_item_id,omitempty" xml:"m_item_id,omitempty"`
}

var poolNrtItemSyncResultDto = sync.Pool{
	New: func() any {
		return new(NrtItemSyncResultDto)
	},
}

// GetNrtItemSyncResultDto() 从对象池中获取NrtItemSyncResultDto
func GetNrtItemSyncResultDto() *NrtItemSyncResultDto {
	return poolNrtItemSyncResultDto.Get().(*NrtItemSyncResultDto)
}

// ReleaseNrtItemSyncResultDto 释放NrtItemSyncResultDto
func ReleaseNrtItemSyncResultDto(v *NrtItemSyncResultDto) {
	v.SItemId = 0
	v.MItemId = 0
	poolNrtItemSyncResultDto.Put(v)
}
