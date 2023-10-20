package ascpchannel

import (
	"sync"
)

// Itemdto 结构体
type Itemdto struct {
	// 前端商品列表
	ItemDOList []Itemdolist `json:"item_d_o_list,omitempty" xml:"item_d_o_list>itemdolist,omitempty"`
	// 前端商品信息
	ItemDoList []Itemdolist `json:"item_do_list,omitempty" xml:"item_do_list>itemdolist,omitempty"`
	// 货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}

var poolItemdto = sync.Pool{
	New: func() any {
		return new(Itemdto)
	},
}

// GetItemdto() 从对象池中获取Itemdto
func GetItemdto() *Itemdto {
	return poolItemdto.Get().(*Itemdto)
}

// ReleaseItemdto 释放Itemdto
func ReleaseItemdto(v *Itemdto) {
	v.ItemDOList = v.ItemDOList[:0]
	v.ItemDoList = v.ItemDoList[:0]
	v.ScItemId = 0
	poolItemdto.Put(v)
}
