package ascp

import (
	"sync"
)

// ItemSkuCost 结构体
type ItemSkuCost struct {
	// 返回错误提示
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// skuId
	Skuid int64 `json:"skuid,omitempty" xml:"skuid,omitempty"`
	// itemId
	Itemid int64 `json:"itemid,omitempty" xml:"itemid,omitempty"`
	// 成本 分为单位
	Cost int64 `json:"cost,omitempty" xml:"cost,omitempty"`
}

var poolItemSkuCost = sync.Pool{
	New: func() any {
		return new(ItemSkuCost)
	},
}

// GetItemSkuCost() 从对象池中获取ItemSkuCost
func GetItemSkuCost() *ItemSkuCost {
	return poolItemSkuCost.Get().(*ItemSkuCost)
}

// ReleaseItemSkuCost 释放ItemSkuCost
func ReleaseItemSkuCost(v *ItemSkuCost) {
	v.Msg = ""
	v.Skuid = 0
	v.Itemid = 0
	v.Cost = 0
	poolItemSkuCost.Put(v)
}
