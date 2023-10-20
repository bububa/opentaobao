package ascpchannel

import (
	"sync"
)

// Outofstockitems 结构体
type Outofstockitems struct {
	// 履约子单号
	SubOrderCode string `json:"sub_order_code,omitempty" xml:"sub_order_code,omitempty"`
	// 货品ID
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 货品缺货数量
	LackQuantity int64 `json:"lack_quantity,omitempty" xml:"lack_quantity,omitempty"`
}

var poolOutofstockitems = sync.Pool{
	New: func() any {
		return new(Outofstockitems)
	},
}

// GetOutofstockitems() 从对象池中获取Outofstockitems
func GetOutofstockitems() *Outofstockitems {
	return poolOutofstockitems.Get().(*Outofstockitems)
}

// ReleaseOutofstockitems 释放Outofstockitems
func ReleaseOutofstockitems(v *Outofstockitems) {
	v.SubOrderCode = ""
	v.ScItemId = ""
	v.LackQuantity = 0
	poolOutofstockitems.Put(v)
}
