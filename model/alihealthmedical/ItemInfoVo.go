package alihealthmedical

import (
	"sync"
)

// ItemInfoVo 结构体
type ItemInfoVo struct {
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolItemInfoVo = sync.Pool{
	New: func() any {
		return new(ItemInfoVo)
	},
}

// GetItemInfoVo() 从对象池中获取ItemInfoVo
func GetItemInfoVo() *ItemInfoVo {
	return poolItemInfoVo.Get().(*ItemInfoVo)
}

// ReleaseItemInfoVo 释放ItemInfoVo
func ReleaseItemInfoVo(v *ItemInfoVo) {
	v.ItemId = 0
	poolItemInfoVo.Put(v)
}
