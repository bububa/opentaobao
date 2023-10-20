package simba

import (
	"sync"
)

// OptionGroupVo 结构体
type OptionGroupVo struct {
	// 标签选项
	OptionList []OptionVo `json:"option_list,omitempty" xml:"option_list>option_vo,omitempty"`
	// 标签分组名称
	OptionGroupName string `json:"option_group_name,omitempty" xml:"option_group_name,omitempty"`
}

var poolOptionGroupVo = sync.Pool{
	New: func() any {
		return new(OptionGroupVo)
	},
}

// GetOptionGroupVo() 从对象池中获取OptionGroupVo
func GetOptionGroupVo() *OptionGroupVo {
	return poolOptionGroupVo.Get().(*OptionGroupVo)
}

// ReleaseOptionGroupVo 释放OptionGroupVo
func ReleaseOptionGroupVo(v *OptionGroupVo) {
	v.OptionList = v.OptionList[:0]
	v.OptionGroupName = ""
	poolOptionGroupVo.Put(v)
}
