package ascpchannel

import (
	"sync"
)

// Tmsitems 结构体
type Tmsitems struct {
	// 履约子单号
	SubOrderCode string `json:"sub_order_code,omitempty" xml:"sub_order_code,omitempty"`
	// 货品ID
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 货品实发数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
	// 货品缺发数量
	LackQuantity int64 `json:"lack_quantity,omitempty" xml:"lack_quantity,omitempty"`
}

var poolTmsitems = sync.Pool{
	New: func() any {
		return new(Tmsitems)
	},
}

// GetTmsitems() 从对象池中获取Tmsitems
func GetTmsitems() *Tmsitems {
	return poolTmsitems.Get().(*Tmsitems)
}

// ReleaseTmsitems 释放Tmsitems
func ReleaseTmsitems(v *Tmsitems) {
	v.SubOrderCode = ""
	v.ScItemId = ""
	v.ItemQuantity = 0
	v.LackQuantity = 0
	poolTmsitems.Put(v)
}
