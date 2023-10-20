package simba

import (
	"sync"
)

// ShowTagVo 结构体
type ShowTagVo struct {
	// 标签code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 展示文案，不为空则展示
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 悬浮描述
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
	// 标签左侧图标地址，不为空则展示
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 背景颜色
	Color string `json:"color,omitempty" xml:"color,omitempty"`
	// 标签类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 扩展属性
	Properties string `json:"properties,omitempty" xml:"properties,omitempty"`
}

var poolShowTagVo = sync.Pool{
	New: func() any {
		return new(ShowTagVo)
	},
}

// GetShowTagVo() 从对象池中获取ShowTagVo
func GetShowTagVo() *ShowTagVo {
	return poolShowTagVo.Get().(*ShowTagVo)
}

// ReleaseShowTagVo 释放ShowTagVo
func ReleaseShowTagVo(v *ShowTagVo) {
	v.Code = ""
	v.Name = ""
	v.Tips = ""
	v.Icon = ""
	v.Color = ""
	v.Type = ""
	v.Properties = ""
	poolShowTagVo.Put(v)
}
