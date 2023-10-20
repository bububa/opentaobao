package simba

import (
	"sync"
)

// CreativeChildrenVo 结构体
type CreativeChildrenVo struct {
	// 素材集合
	MaterialList []ItemMaterialVo `json:"material_list,omitempty" xml:"material_list>item_material_vo,omitempty"`
}

var poolCreativeChildrenVo = sync.Pool{
	New: func() any {
		return new(CreativeChildrenVo)
	},
}

// GetCreativeChildrenVo() 从对象池中获取CreativeChildrenVo
func GetCreativeChildrenVo() *CreativeChildrenVo {
	return poolCreativeChildrenVo.Get().(*CreativeChildrenVo)
}

// ReleaseCreativeChildrenVo 释放CreativeChildrenVo
func ReleaseCreativeChildrenVo(v *CreativeChildrenVo) {
	v.MaterialList = v.MaterialList[:0]
	poolCreativeChildrenVo.Put(v)
}
