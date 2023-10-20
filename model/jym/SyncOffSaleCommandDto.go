package jym

import (
	"sync"
)

// SyncOffSaleCommandDto 结构体
type SyncOffSaleCommandDto struct {
	// 交易猫商品ID
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
}

var poolSyncOffSaleCommandDto = sync.Pool{
	New: func() any {
		return new(SyncOffSaleCommandDto)
	},
}

// GetSyncOffSaleCommandDto() 从对象池中获取SyncOffSaleCommandDto
func GetSyncOffSaleCommandDto() *SyncOffSaleCommandDto {
	return poolSyncOffSaleCommandDto.Get().(*SyncOffSaleCommandDto)
}

// ReleaseSyncOffSaleCommandDto 释放SyncOffSaleCommandDto
func ReleaseSyncOffSaleCommandDto(v *SyncOffSaleCommandDto) {
	v.GoodsId = 0
	poolSyncOffSaleCommandDto.Put(v)
}
