package trade

import (
	"sync"
)

// AdvertiseInfoDto 结构体
type AdvertiseInfoDto struct {
	// 信息流投放转换追踪标识
	ConversionTracking string `json:"conversion_tracking,omitempty" xml:"conversion_tracking,omitempty"`
	// 信息流投放广告账户id
	AdvertiserId string `json:"advertiser_id,omitempty" xml:"advertiser_id,omitempty"`
}

var poolAdvertiseInfoDto = sync.Pool{
	New: func() any {
		return new(AdvertiseInfoDto)
	},
}

// GetAdvertiseInfoDto() 从对象池中获取AdvertiseInfoDto
func GetAdvertiseInfoDto() *AdvertiseInfoDto {
	return poolAdvertiseInfoDto.Get().(*AdvertiseInfoDto)
}

// ReleaseAdvertiseInfoDto 释放AdvertiseInfoDto
func ReleaseAdvertiseInfoDto(v *AdvertiseInfoDto) {
	v.ConversionTracking = ""
	v.AdvertiserId = ""
	poolAdvertiseInfoDto.Put(v)
}
