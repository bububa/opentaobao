package servicecenter

import (
	"sync"
)

// OfnRecycleOrderSyncOfflineSettleInfoDto 结构体
type OfnRecycleOrderSyncOfflineSettleInfoDto struct {
	// 是否同步成功
	SyncSuccess bool `json:"sync_success,omitempty" xml:"sync_success,omitempty"`
}

var poolOfnRecycleOrderSyncOfflineSettleInfoDto = sync.Pool{
	New: func() any {
		return new(OfnRecycleOrderSyncOfflineSettleInfoDto)
	},
}

// GetOfnRecycleOrderSyncOfflineSettleInfoDto() 从对象池中获取OfnRecycleOrderSyncOfflineSettleInfoDto
func GetOfnRecycleOrderSyncOfflineSettleInfoDto() *OfnRecycleOrderSyncOfflineSettleInfoDto {
	return poolOfnRecycleOrderSyncOfflineSettleInfoDto.Get().(*OfnRecycleOrderSyncOfflineSettleInfoDto)
}

// ReleaseOfnRecycleOrderSyncOfflineSettleInfoDto 释放OfnRecycleOrderSyncOfflineSettleInfoDto
func ReleaseOfnRecycleOrderSyncOfflineSettleInfoDto(v *OfnRecycleOrderSyncOfflineSettleInfoDto) {
	v.SyncSuccess = false
	poolOfnRecycleOrderSyncOfflineSettleInfoDto.Put(v)
}
