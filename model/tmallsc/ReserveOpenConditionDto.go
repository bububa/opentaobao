package tmallsc

import (
	"sync"
)

// ReserveOpenConditionDto 结构体
type ReserveOpenConditionDto struct {
	// 服务code
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 排除的区域id
	ExcludeAreaIds string `json:"exclude_area_ids,omitempty" xml:"exclude_area_ids,omitempty"`
	// 区域ids集合
	AreaIds string `json:"area_ids,omitempty" xml:"area_ids,omitempty"`
	// 类目id
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 品牌id
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 城市id
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 身份id
	ProvinceId int64 `json:"province_id,omitempty" xml:"province_id,omitempty"`
}

var poolReserveOpenConditionDto = sync.Pool{
	New: func() any {
		return new(ReserveOpenConditionDto)
	},
}

// GetReserveOpenConditionDto() 从对象池中获取ReserveOpenConditionDto
func GetReserveOpenConditionDto() *ReserveOpenConditionDto {
	return poolReserveOpenConditionDto.Get().(*ReserveOpenConditionDto)
}

// ReleaseReserveOpenConditionDto 释放ReserveOpenConditionDto
func ReleaseReserveOpenConditionDto(v *ReserveOpenConditionDto) {
	v.ServiceCode = ""
	v.ExcludeAreaIds = ""
	v.AreaIds = ""
	v.CategoryId = 0
	v.BrandId = 0
	v.CityId = 0
	v.ProvinceId = 0
	poolReserveOpenConditionDto.Put(v)
}
