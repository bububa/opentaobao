package alitripmerchant

import (
	"sync"
)

// FacilityListVo 结构体
type FacilityListVo struct {
	// 设施集合
	FacilityList []FacilityVo `json:"facility_list,omitempty" xml:"facility_list>facility_vo,omitempty"`
	// 设施分组名
	FacilityName string `json:"facility_name,omitempty" xml:"facility_name,omitempty"`
}

var poolFacilityListVo = sync.Pool{
	New: func() any {
		return new(FacilityListVo)
	},
}

// GetFacilityListVo() 从对象池中获取FacilityListVo
func GetFacilityListVo() *FacilityListVo {
	return poolFacilityListVo.Get().(*FacilityListVo)
}

// ReleaseFacilityListVo 释放FacilityListVo
func ReleaseFacilityListVo(v *FacilityListVo) {
	v.FacilityList = v.FacilityList[:0]
	v.FacilityName = ""
	poolFacilityListVo.Put(v)
}
