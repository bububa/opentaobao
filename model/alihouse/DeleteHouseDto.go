package alihouse

import (
	"sync"
)

// DeleteHouseDto 结构体
type DeleteHouseDto struct {
	// 外部小区id
	CommunityOuterId string `json:"community_outer_id,omitempty" xml:"community_outer_id,omitempty"`
	// 外部房源id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 委托房源外部房源id
	EntrustOuterId string `json:"entrust_outer_id,omitempty" xml:"entrust_outer_id,omitempty"`
	// 业务类型，1-新房，2-二手房，3-租房，默认为2
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 0-正常房源 1-临时房源（交易专用）
	HouseType int64 `json:"house_type,omitempty" xml:"house_type,omitempty"`
	// 租房业务模式
	HouseModel int64 `json:"house_model,omitempty" xml:"house_model,omitempty"`
	// 公私域区分 1公域 2私域
	Scene int64 `json:"scene,omitempty" xml:"scene,omitempty"`
}

var poolDeleteHouseDto = sync.Pool{
	New: func() any {
		return new(DeleteHouseDto)
	},
}

// GetDeleteHouseDto() 从对象池中获取DeleteHouseDto
func GetDeleteHouseDto() *DeleteHouseDto {
	return poolDeleteHouseDto.Get().(*DeleteHouseDto)
}

// ReleaseDeleteHouseDto 释放DeleteHouseDto
func ReleaseDeleteHouseDto(v *DeleteHouseDto) {
	v.CommunityOuterId = ""
	v.OuterId = ""
	v.EntrustOuterId = ""
	v.BusinessType = 0
	v.HouseType = 0
	v.HouseModel = 0
	v.Scene = 0
	poolDeleteHouseDto.Put(v)
}
