package alitripmerchant

import (
	"sync"
)

// FacilityVo 结构体
type FacilityVo struct {
	// 图片icon
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 设施名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 概要
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 设施Id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolFacilityVo = sync.Pool{
	New: func() any {
		return new(FacilityVo)
	},
}

// GetFacilityVo() 从对象池中获取FacilityVo
func GetFacilityVo() *FacilityVo {
	return poolFacilityVo.Get().(*FacilityVo)
}

// ReleaseFacilityVo 释放FacilityVo
func ReleaseFacilityVo(v *FacilityVo) {
	v.Icon = ""
	v.Name = ""
	v.Summary = ""
	v.Description = ""
	v.Code = ""
	v.Id = 0
	poolFacilityVo.Put(v)
}
