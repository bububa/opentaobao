package logistic

import (
	"sync"
)

// GoodsItem 结构体
type GoodsItem struct {
	// 货品ID
	GoodsId string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 子订单编号
	Oid int64 `json:"oid,omitempty" xml:"oid,omitempty"`
	// 货品数量
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}

var poolGoodsItem = sync.Pool{
	New: func() any {
		return new(GoodsItem)
	},
}

// GetGoodsItem() 从对象池中获取GoodsItem
func GetGoodsItem() *GoodsItem {
	return poolGoodsItem.Get().(*GoodsItem)
}

// ReleaseGoodsItem 释放GoodsItem
func ReleaseGoodsItem(v *GoodsItem) {
	v.GoodsId = ""
	v.Oid = 0
	v.Num = 0
	poolGoodsItem.Put(v)
}
