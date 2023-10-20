package tmallsc

import (
	"sync"
)

// ReserveOpenConditionDelDto 结构体
type ReserveOpenConditionDelDto struct {
	// 服务code
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 区域集合ids
	AreaIds string `json:"area_ids,omitempty" xml:"area_ids,omitempty"`
	// 城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 品牌id
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 省份id
	ProvinceId int64 `json:"province_id,omitempty" xml:"province_id,omitempty"`
}

var poolReserveOpenConditionDelDto = sync.Pool{
	New: func() any {
		return new(ReserveOpenConditionDelDto)
	},
}

// GetReserveOpenConditionDelDto() 从对象池中获取ReserveOpenConditionDelDto
func GetReserveOpenConditionDelDto() *ReserveOpenConditionDelDto {
	return poolReserveOpenConditionDelDto.Get().(*ReserveOpenConditionDelDto)
}

// ReleaseReserveOpenConditionDelDto 释放ReserveOpenConditionDelDto
func ReleaseReserveOpenConditionDelDto(v *ReserveOpenConditionDelDto) {
	v.ServiceCode = ""
	v.AreaIds = ""
	v.CityId = 0
	v.CategoryId = 0
	v.BrandId = 0
	v.ProvinceId = 0
	poolReserveOpenConditionDelDto.Put(v)
}
