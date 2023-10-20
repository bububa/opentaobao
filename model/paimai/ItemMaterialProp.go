package paimai

import (
	"sync"
)

// ItemMaterialProp 结构体
type ItemMaterialProp struct {
	// 材质值列表
	Materials []ItemMateriaValueDo `json:"materials,omitempty" xml:"materials>item_materia_value_do,omitempty"`
}

var poolItemMaterialProp = sync.Pool{
	New: func() any {
		return new(ItemMaterialProp)
	},
}

// GetItemMaterialProp() 从对象池中获取ItemMaterialProp
func GetItemMaterialProp() *ItemMaterialProp {
	return poolItemMaterialProp.Get().(*ItemMaterialProp)
}

// ReleaseItemMaterialProp 释放ItemMaterialProp
func ReleaseItemMaterialProp(v *ItemMaterialProp) {
	v.Materials = v.Materials[:0]
	poolItemMaterialProp.Put(v)
}
