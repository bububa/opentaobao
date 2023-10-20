package servicecenter

import (
	"sync"
)

// OfnRecycleOrderSyncFulfillTypeDto 结构体
type OfnRecycleOrderSyncFulfillTypeDto struct {
	// 是否同步成功
	SyncSuccess bool `json:"sync_success,omitempty" xml:"sync_success,omitempty"`
}

var poolOfnRecycleOrderSyncFulfillTypeDto = sync.Pool{
	New: func() any {
		return new(OfnRecycleOrderSyncFulfillTypeDto)
	},
}

// GetOfnRecycleOrderSyncFulfillTypeDto() 从对象池中获取OfnRecycleOrderSyncFulfillTypeDto
func GetOfnRecycleOrderSyncFulfillTypeDto() *OfnRecycleOrderSyncFulfillTypeDto {
	return poolOfnRecycleOrderSyncFulfillTypeDto.Get().(*OfnRecycleOrderSyncFulfillTypeDto)
}

// ReleaseOfnRecycleOrderSyncFulfillTypeDto 释放OfnRecycleOrderSyncFulfillTypeDto
func ReleaseOfnRecycleOrderSyncFulfillTypeDto(v *OfnRecycleOrderSyncFulfillTypeDto) {
	v.SyncSuccess = false
	poolOfnRecycleOrderSyncFulfillTypeDto.Put(v)
}
