package ascpchannel

import (
	"sync"
)

// WaybillGenItemList 结构体
type WaybillGenItemList struct {
	// 货品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货品数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
}

var poolWaybillGenItemList = sync.Pool{
	New: func() any {
		return new(WaybillGenItemList)
	},
}

// GetWaybillGenItemList() 从对象池中获取WaybillGenItemList
func GetWaybillGenItemList() *WaybillGenItemList {
	return poolWaybillGenItemList.Get().(*WaybillGenItemList)
}

// ReleaseWaybillGenItemList 释放WaybillGenItemList
func ReleaseWaybillGenItemList(v *WaybillGenItemList) {
	v.ItemId = ""
	v.ItemQuantity = 0
	poolWaybillGenItemList.Put(v)
}
