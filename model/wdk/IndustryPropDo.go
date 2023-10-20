package wdk

import (
	"sync"
)

// IndustryPropDo 结构体
type IndustryPropDo struct {
	// 行业对应的属性
	Props []PropDo `json:"props,omitempty" xml:"props>prop_do,omitempty"`
	// 行业类型
	IndustryType string `json:"industry_type,omitempty" xml:"industry_type,omitempty"`
}

var poolIndustryPropDo = sync.Pool{
	New: func() any {
		return new(IndustryPropDo)
	},
}

// GetIndustryPropDo() 从对象池中获取IndustryPropDo
func GetIndustryPropDo() *IndustryPropDo {
	return poolIndustryPropDo.Get().(*IndustryPropDo)
}

// ReleaseIndustryPropDo 释放IndustryPropDo
func ReleaseIndustryPropDo(v *IndustryPropDo) {
	v.Props = v.Props[:0]
	v.IndustryType = ""
	poolIndustryPropDo.Put(v)
}
