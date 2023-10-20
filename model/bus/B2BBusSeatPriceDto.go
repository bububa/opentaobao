package bus

import (
	"sync"
)

// B2BBusSeatPriceDto 结构体
type B2BBusSeatPriceDto struct {
	// 出发城市
	DepCityName string `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	// 出发时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 目的地
	LastPlaceName string `json:"last_place_name,omitempty" xml:"last_place_name,omitempty"`
	// 车次ID
	ScheduleId string `json:"schedule_id,omitempty" xml:"schedule_id,omitempty"`
	// 车次全价
	FullPrice int64 `json:"full_price,omitempty" xml:"full_price,omitempty"`
	// 余票数量
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
	// 服务费
	ServicePrice int64 `json:"service_price,omitempty" xml:"service_price,omitempty"`
}

var poolB2BBusSeatPriceDto = sync.Pool{
	New: func() any {
		return new(B2BBusSeatPriceDto)
	},
}

// GetB2BBusSeatPriceDto() 从对象池中获取B2BBusSeatPriceDto
func GetB2BBusSeatPriceDto() *B2BBusSeatPriceDto {
	return poolB2BBusSeatPriceDto.Get().(*B2BBusSeatPriceDto)
}

// ReleaseB2BBusSeatPriceDto 释放B2BBusSeatPriceDto
func ReleaseB2BBusSeatPriceDto(v *B2BBusSeatPriceDto) {
	v.DepCityName = ""
	v.DepTime = ""
	v.LastPlaceName = ""
	v.ScheduleId = ""
	v.FullPrice = 0
	v.Stock = 0
	v.ServicePrice = 0
	poolB2BBusSeatPriceDto.Put(v)
}
