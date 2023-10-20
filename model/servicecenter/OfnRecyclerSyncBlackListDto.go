package servicecenter

import (
	"sync"
)

// OfnRecyclerSyncBlackListDto 结构体
type OfnRecyclerSyncBlackListDto struct {
	// 是否同步成功
	SyncSuccess bool `json:"sync_success,omitempty" xml:"sync_success,omitempty"`
}

var poolOfnRecyclerSyncBlackListDto = sync.Pool{
	New: func() any {
		return new(OfnRecyclerSyncBlackListDto)
	},
}

// GetOfnRecyclerSyncBlackListDto() 从对象池中获取OfnRecyclerSyncBlackListDto
func GetOfnRecyclerSyncBlackListDto() *OfnRecyclerSyncBlackListDto {
	return poolOfnRecyclerSyncBlackListDto.Get().(*OfnRecyclerSyncBlackListDto)
}

// ReleaseOfnRecyclerSyncBlackListDto 释放OfnRecyclerSyncBlackListDto
func ReleaseOfnRecyclerSyncBlackListDto(v *OfnRecyclerSyncBlackListDto) {
	v.SyncSuccess = false
	poolOfnRecyclerSyncBlackListDto.Put(v)
}
