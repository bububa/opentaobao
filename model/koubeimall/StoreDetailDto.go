package koubeimall

import (
	"sync"
)

// StoreDetailDto 结构体
type StoreDetailDto struct {
	// 门店相册
	StoreAlbum *StoreAlbumDto `json:"store_album,omitempty" xml:"store_album,omitempty"`
	// 门店基础信息
	StoreDto *StoreDto `json:"store_dto,omitempty" xml:"store_dto,omitempty"`
	// 服务信息
	ServiceInfo *ServiceInfoDto `json:"service_info,omitempty" xml:"service_info,omitempty"`
}

var poolStoreDetailDto = sync.Pool{
	New: func() any {
		return new(StoreDetailDto)
	},
}

// GetStoreDetailDto() 从对象池中获取StoreDetailDto
func GetStoreDetailDto() *StoreDetailDto {
	return poolStoreDetailDto.Get().(*StoreDetailDto)
}

// ReleaseStoreDetailDto 释放StoreDetailDto
func ReleaseStoreDetailDto(v *StoreDetailDto) {
	v.StoreAlbum = nil
	v.StoreDto = nil
	v.ServiceInfo = nil
	poolStoreDetailDto.Put(v)
}
