package paimai

import (
	"sync"
)

// ItemMateriaValueDo 结构体
type ItemMateriaValueDo struct {
	// 材质值名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 当前材质值，是否需要填写含量值。比如：棉 是需要填写含量值，而牛皮 是不需要填写含量值的
	NeedContentNumber bool `json:"need_content_number,omitempty" xml:"need_content_number,omitempty"`
}

var poolItemMateriaValueDo = sync.Pool{
	New: func() any {
		return new(ItemMateriaValueDo)
	},
}

// GetItemMateriaValueDo() 从对象池中获取ItemMateriaValueDo
func GetItemMateriaValueDo() *ItemMateriaValueDo {
	return poolItemMateriaValueDo.Get().(*ItemMateriaValueDo)
}

// ReleaseItemMateriaValueDo 释放ItemMateriaValueDo
func ReleaseItemMateriaValueDo(v *ItemMateriaValueDo) {
	v.Name = ""
	v.NeedContentNumber = false
	poolItemMateriaValueDo.Put(v)
}
