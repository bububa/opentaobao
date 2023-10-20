package simba

import (
	"sync"
)

// PreAddItemCreativeVo 结构体
type PreAddItemCreativeVo struct {
	// 单品创意素材
	Material *ItemMaterialVo `json:"material,omitempty" xml:"material,omitempty"`
	// 是否有长素材
	HaveLongSucai bool `json:"have_long_sucai,omitempty" xml:"have_long_sucai,omitempty"`
}

var poolPreAddItemCreativeVo = sync.Pool{
	New: func() any {
		return new(PreAddItemCreativeVo)
	},
}

// GetPreAddItemCreativeVo() 从对象池中获取PreAddItemCreativeVo
func GetPreAddItemCreativeVo() *PreAddItemCreativeVo {
	return poolPreAddItemCreativeVo.Get().(*PreAddItemCreativeVo)
}

// ReleasePreAddItemCreativeVo 释放PreAddItemCreativeVo
func ReleasePreAddItemCreativeVo(v *PreAddItemCreativeVo) {
	v.Material = nil
	v.HaveLongSucai = false
	poolPreAddItemCreativeVo.Put(v)
}
