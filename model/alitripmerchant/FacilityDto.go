package alitripmerchant

import (
	"sync"
)

// FacilityDto 结构体
type FacilityDto struct {
	// 设施图片
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 设施名字
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 设施id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolFacilityDto = sync.Pool{
	New: func() any {
		return new(FacilityDto)
	},
}

// GetFacilityDto() 从对象池中获取FacilityDto
func GetFacilityDto() *FacilityDto {
	return poolFacilityDto.Get().(*FacilityDto)
}

// ReleaseFacilityDto 释放FacilityDto
func ReleaseFacilityDto(v *FacilityDto) {
	v.Icon = ""
	v.Name = ""
	v.Id = 0
	poolFacilityDto.Put(v)
}
