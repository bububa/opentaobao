package aedropshiper

import (
	"sync"
)

// LogisticsInfoDto 结构体
type LogisticsInfoDto struct {
	// Country
	ShipToCountry string `json:"ship_to_country,omitempty" xml:"ship_to_country,omitempty"`
	// Goods lead time
	DeliveryTime int64 `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
}

var poolLogisticsInfoDto = sync.Pool{
	New: func() any {
		return new(LogisticsInfoDto)
	},
}

// GetLogisticsInfoDto() 从对象池中获取LogisticsInfoDto
func GetLogisticsInfoDto() *LogisticsInfoDto {
	return poolLogisticsInfoDto.Get().(*LogisticsInfoDto)
}

// ReleaseLogisticsInfoDto 释放LogisticsInfoDto
func ReleaseLogisticsInfoDto(v *LogisticsInfoDto) {
	v.ShipToCountry = ""
	v.DeliveryTime = 0
	poolLogisticsInfoDto.Put(v)
}
