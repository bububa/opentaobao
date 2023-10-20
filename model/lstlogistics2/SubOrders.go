package lstlogistics2

import (
	"sync"
)

// SubOrders 结构体
type SubOrders struct {
	// 外部商品编码
	OutItemCode string `json:"out_item_code,omitempty" xml:"out_item_code,omitempty"`
	// 商品条码
	ItemBarCode string `json:"item_bar_code,omitempty" xml:"item_bar_code,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 零售通子订单号
	LstSubOrderId string `json:"lst_sub_order_id,omitempty" xml:"lst_sub_order_id,omitempty"`
	// 外部子订单号
	OutSubOrderId string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// 商品数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
	// 揽收数量
	PickQuantity int64 `json:"pick_quantity,omitempty" xml:"pick_quantity,omitempty"`
	// 签收数量
	SignQuantity int64 `json:"sign_quantity,omitempty" xml:"sign_quantity,omitempty"`
}

var poolSubOrders = sync.Pool{
	New: func() any {
		return new(SubOrders)
	},
}

// GetSubOrders() 从对象池中获取SubOrders
func GetSubOrders() *SubOrders {
	return poolSubOrders.Get().(*SubOrders)
}

// ReleaseSubOrders 释放SubOrders
func ReleaseSubOrders(v *SubOrders) {
	v.OutItemCode = ""
	v.ItemBarCode = ""
	v.ItemName = ""
	v.LstSubOrderId = ""
	v.OutSubOrderId = ""
	v.ItemQuantity = 0
	v.PickQuantity = 0
	v.SignQuantity = 0
	poolSubOrders.Put(v)
}
