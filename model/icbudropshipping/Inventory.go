package icbudropshipping

import (
	"sync"
)

// Inventory 结构体
type Inventory struct {
	// dispatch country
	DispatchCountry string `json:"dispatch_country,omitempty" xml:"dispatch_country,omitempty"`
	// store code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// inventory count
	Inventory int64 `json:"inventory,omitempty" xml:"inventory,omitempty"`
}

var poolInventory = sync.Pool{
	New: func() any {
		return new(Inventory)
	},
}

// GetInventory() 从对象池中获取Inventory
func GetInventory() *Inventory {
	return poolInventory.Get().(*Inventory)
}

// ReleaseInventory 释放Inventory
func ReleaseInventory(v *Inventory) {
	v.DispatchCountry = ""
	v.StoreCode = ""
	v.Inventory = 0
	poolInventory.Put(v)
}
