package tmallcar

import (
	"sync"
)

// StoreServiceItem4IsvDto 结构体
type StoreServiceItem4IsvDto struct {
	// 商品列表
	StoreServiceItems []Item4IsvDto `json:"store_service_items,omitempty" xml:"store_service_items>item4isv_dto,omitempty"`
}

var poolStoreServiceItem4IsvDto = sync.Pool{
	New: func() any {
		return new(StoreServiceItem4IsvDto)
	},
}

// GetStoreServiceItem4IsvDto() 从对象池中获取StoreServiceItem4IsvDto
func GetStoreServiceItem4IsvDto() *StoreServiceItem4IsvDto {
	return poolStoreServiceItem4IsvDto.Get().(*StoreServiceItem4IsvDto)
}

// ReleaseStoreServiceItem4IsvDto 释放StoreServiceItem4IsvDto
func ReleaseStoreServiceItem4IsvDto(v *StoreServiceItem4IsvDto) {
	v.StoreServiceItems = v.StoreServiceItems[:0]
	poolStoreServiceItem4IsvDto.Put(v)
}
