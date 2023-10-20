package alihouse

import (
	"sync"
)

// HouseFeaturesDto 结构体
type HouseFeaturesDto struct {
	// 外部房源id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部小区id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 租房业务模式
	HouseModel int64 `json:"house_model,omitempty" xml:"house_model,omitempty"`
}

var poolHouseFeaturesDto = sync.Pool{
	New: func() any {
		return new(HouseFeaturesDto)
	},
}

// GetHouseFeaturesDto() 从对象池中获取HouseFeaturesDto
func GetHouseFeaturesDto() *HouseFeaturesDto {
	return poolHouseFeaturesDto.Get().(*HouseFeaturesDto)
}

// ReleaseHouseFeaturesDto 释放HouseFeaturesDto
func ReleaseHouseFeaturesDto(v *HouseFeaturesDto) {
	v.OuterId = ""
	v.CommunityOuterId = ""
	v.HouseModel = 0
	poolHouseFeaturesDto.Put(v)
}
