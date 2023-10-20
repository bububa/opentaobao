package tmic

import (
	"sync"
)

// ItemBo 结构体
type ItemBo struct {
	// 选项所对应的图片cdn地址
	Img string `json:"img,omitempty" xml:"img,omitempty"`
	// 该选项的唯一编码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 该选项的说明
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

var poolItemBo = sync.Pool{
	New: func() any {
		return new(ItemBo)
	},
}

// GetItemBo() 从对象池中获取ItemBo
func GetItemBo() *ItemBo {
	return poolItemBo.Get().(*ItemBo)
}

// ReleaseItemBo 释放ItemBo
func ReleaseItemBo(v *ItemBo) {
	v.Img = ""
	v.Code = ""
	v.Value = ""
	poolItemBo.Put(v)
}
