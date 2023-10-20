package alihealthoutflow

import (
	"sync"
)

// ItemInStockCardVo 结构体
type ItemInStockCardVo struct {
	// 疫苗卡片数据
	ItemInStockList []ItemInStockVo `json:"item_in_stock_list,omitempty" xml:"item_in_stock_list>item_in_stock_vo,omitempty"`
	// 标签名
	TabName string `json:"tab_name,omitempty" xml:"tab_name,omitempty"`
}

var poolItemInStockCardVo = sync.Pool{
	New: func() any {
		return new(ItemInStockCardVo)
	},
}

// GetItemInStockCardVo() 从对象池中获取ItemInStockCardVo
func GetItemInStockCardVo() *ItemInStockCardVo {
	return poolItemInStockCardVo.Get().(*ItemInStockCardVo)
}

// ReleaseItemInStockCardVo 释放ItemInStockCardVo
func ReleaseItemInStockCardVo(v *ItemInStockCardVo) {
	v.ItemInStockList = v.ItemInStockList[:0]
	v.TabName = ""
	poolItemInStockCardVo.Put(v)
}
