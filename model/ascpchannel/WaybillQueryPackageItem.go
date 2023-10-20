package ascpchannel

import (
	"sync"
)

// WaybillQueryPackageItem 结构体
type WaybillQueryPackageItem struct {
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品code
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 商品数量
	ItemCount string `json:"item_count,omitempty" xml:"item_count,omitempty"`
}

var poolWaybillQueryPackageItem = sync.Pool{
	New: func() any {
		return new(WaybillQueryPackageItem)
	},
}

// GetWaybillQueryPackageItem() 从对象池中获取WaybillQueryPackageItem
func GetWaybillQueryPackageItem() *WaybillQueryPackageItem {
	return poolWaybillQueryPackageItem.Get().(*WaybillQueryPackageItem)
}

// ReleaseWaybillQueryPackageItem 释放WaybillQueryPackageItem
func ReleaseWaybillQueryPackageItem(v *WaybillQueryPackageItem) {
	v.ItemName = ""
	v.ItemCode = ""
	v.ItemCount = ""
	poolWaybillQueryPackageItem.Put(v)
}
