package tbk

import (
	"sync"
)

// Ucrowdrankitems 结构体
type Ucrowdrankitems struct {
	// 物料评估-商品ID，material_id=41377时必填
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 物料评估-商品佣金率，如：1234表示12.34%，material_id=41377时选填
	Commirate int64 `json:"commirate,omitempty" xml:"commirate,omitempty"`
	// 物料评估-商品价格，单位：元，material_id=41377时选填
	Price float64 `json:"price,omitempty" xml:"price,omitempty"`
}

var poolUcrowdrankitems = sync.Pool{
	New: func() any {
		return new(Ucrowdrankitems)
	},
}

// GetUcrowdrankitems() 从对象池中获取Ucrowdrankitems
func GetUcrowdrankitems() *Ucrowdrankitems {
	return poolUcrowdrankitems.Get().(*Ucrowdrankitems)
}

// ReleaseUcrowdrankitems 释放Ucrowdrankitems
func ReleaseUcrowdrankitems(v *Ucrowdrankitems) {
	v.ItemId = ""
	v.Commirate = 0
	v.Price = 0
	poolUcrowdrankitems.Put(v)
}
