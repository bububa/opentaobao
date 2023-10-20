package alihouse

import (
	"sync"
)

// GoodsHouseDto 结构体
type GoodsHouseDto struct {
	// 委托房源所属小区外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 房源外部id
	OuterHouseId string `json:"outer_house_id,omitempty" xml:"outer_house_id,omitempty"`
	// 房源业务类型 2-分散房源 3 -集中式房源
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
}

var poolGoodsHouseDto = sync.Pool{
	New: func() any {
		return new(GoodsHouseDto)
	},
}

// GetGoodsHouseDto() 从对象池中获取GoodsHouseDto
func GetGoodsHouseDto() *GoodsHouseDto {
	return poolGoodsHouseDto.Get().(*GoodsHouseDto)
}

// ReleaseGoodsHouseDto 释放GoodsHouseDto
func ReleaseGoodsHouseDto(v *GoodsHouseDto) {
	v.OuterId = ""
	v.OuterHouseId = ""
	v.BusinessType = 0
	poolGoodsHouseDto.Put(v)
}
