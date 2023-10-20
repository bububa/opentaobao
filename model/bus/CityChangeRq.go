package bus

import (
	"sync"
)

// CityChangeRq 结构体
type CityChangeRq struct {
	// 城市列表属性
	List []BusCityChangeDto `json:"list,omitempty" xml:"list>bus_city_change_dto,omitempty"`
	// 代理商ID
	AgentId int64 `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
}

var poolCityChangeRq = sync.Pool{
	New: func() any {
		return new(CityChangeRq)
	},
}

// GetCityChangeRq() 从对象池中获取CityChangeRq
func GetCityChangeRq() *CityChangeRq {
	return poolCityChangeRq.Get().(*CityChangeRq)
}

// ReleaseCityChangeRq 释放CityChangeRq
func ReleaseCityChangeRq(v *CityChangeRq) {
	v.List = v.List[:0]
	v.AgentId = 0
	poolCityChangeRq.Put(v)
}
