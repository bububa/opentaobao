package tmallnr

import (
	"sync"
)

// NrtStoreDto 结构体
type NrtStoreDto struct {
	// 门店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 地址
	ShortAddress string `json:"short_address,omitempty" xml:"short_address,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolNrtStoreDto = sync.Pool{
	New: func() any {
		return new(NrtStoreDto)
	},
}

// GetNrtStoreDto() 从对象池中获取NrtStoreDto
func GetNrtStoreDto() *NrtStoreDto {
	return poolNrtStoreDto.Get().(*NrtStoreDto)
}

// ReleaseNrtStoreDto 释放NrtStoreDto
func ReleaseNrtStoreDto(v *NrtStoreDto) {
	v.Name = ""
	v.ShortAddress = ""
	v.Lat = ""
	v.Lng = ""
	v.StoreId = 0
	poolNrtStoreDto.Put(v)
}
