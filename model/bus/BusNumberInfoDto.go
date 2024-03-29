package bus

import (
	"sync"
)

// BusNumberInfoDto 结构体
type BusNumberInfoDto struct {
	// 车次id
	ScheduleId string `json:"schedule_id,omitempty" xml:"schedule_id,omitempty"`
	// 出发时间
	DepartTime string `json:"depart_time,omitempty" xml:"depart_time,omitempty"`
	// 出发城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 代理商id
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	// 库存
	Stock int64 `json:"stock,omitempty" xml:"stock,omitempty"`
	// 价格
	FullPrice int64 `json:"full_price,omitempty" xml:"full_price,omitempty"`
}

var poolBusNumberInfoDto = sync.Pool{
	New: func() any {
		return new(BusNumberInfoDto)
	},
}

// GetBusNumberInfoDto() 从对象池中获取BusNumberInfoDto
func GetBusNumberInfoDto() *BusNumberInfoDto {
	return poolBusNumberInfoDto.Get().(*BusNumberInfoDto)
}

// ReleaseBusNumberInfoDto 释放BusNumberInfoDto
func ReleaseBusNumberInfoDto(v *BusNumberInfoDto) {
	v.ScheduleId = ""
	v.DepartTime = ""
	v.CityName = ""
	v.AgentId = 0
	v.Stock = 0
	v.FullPrice = 0
	poolBusNumberInfoDto.Put(v)
}
