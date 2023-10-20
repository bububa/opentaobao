package alihouse

import (
	"sync"
)

// StoreStatusDto 结构体
type StoreStatusDto struct {
	// 门店外部id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 门店状态 1-正常 2-暂停 3-关闭
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}

var poolStoreStatusDto = sync.Pool{
	New: func() any {
		return new(StoreStatusDto)
	},
}

// GetStoreStatusDto() 从对象池中获取StoreStatusDto
func GetStoreStatusDto() *StoreStatusDto {
	return poolStoreStatusDto.Get().(*StoreStatusDto)
}

// ReleaseStoreStatusDto 释放StoreStatusDto
func ReleaseStoreStatusDto(v *StoreStatusDto) {
	v.OuterStoreId = ""
	v.Status = 0
	poolStoreStatusDto.Put(v)
}
