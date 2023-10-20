package tmic

import (
	"sync"
)

// OptionBo 结构体
type OptionBo struct {
	// 选项具体值
	Items []ItemBo `json:"items,omitempty" xml:"items>item_bo,omitempty"`
	// optionItemBOList
	OptionItemBoList []OptionItemBo `json:"option_item_bo_list,omitempty" xml:"option_item_bo_list>option_item_bo,omitempty"`
	// 是否还有其他选项
	HasOther bool `json:"has_other,omitempty" xml:"has_other,omitempty"`
}

var poolOptionBo = sync.Pool{
	New: func() any {
		return new(OptionBo)
	},
}

// GetOptionBo() 从对象池中获取OptionBo
func GetOptionBo() *OptionBo {
	return poolOptionBo.Get().(*OptionBo)
}

// ReleaseOptionBo 释放OptionBo
func ReleaseOptionBo(v *OptionBo) {
	v.Items = v.Items[:0]
	v.OptionItemBoList = v.OptionItemBoList[:0]
	v.HasOther = false
	poolOptionBo.Put(v)
}
