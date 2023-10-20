package wdk

import (
	"sync"
)

// Inventoryinfolist 结构体
type Inventoryinfolist struct {
	// realInvent
	RealInvent string `json:"real_invent,omitempty" xml:"real_invent,omitempty"`
	// storageUnit
	StorageUnit string `json:"storage_unit,omitempty" xml:"storage_unit,omitempty"`
	// itemCode
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
}

var poolInventoryinfolist = sync.Pool{
	New: func() any {
		return new(Inventoryinfolist)
	},
}

// GetInventoryinfolist() 从对象池中获取Inventoryinfolist
func GetInventoryinfolist() *Inventoryinfolist {
	return poolInventoryinfolist.Get().(*Inventoryinfolist)
}

// ReleaseInventoryinfolist 释放Inventoryinfolist
func ReleaseInventoryinfolist(v *Inventoryinfolist) {
	v.RealInvent = ""
	v.StorageUnit = ""
	v.ItemCode = ""
	poolInventoryinfolist.Put(v)
}
