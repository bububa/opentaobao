package ascpchannel

import (
	"sync"
)

// Packageitemsigninfodtolist 结构体
type Packageitemsigninfodtolist struct {
	// 签收状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 商品数量
	ItemQuantity int64 `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
}

var poolPackageitemsigninfodtolist = sync.Pool{
	New: func() any {
		return new(Packageitemsigninfodtolist)
	},
}

// GetPackageitemsigninfodtolist() 从对象池中获取Packageitemsigninfodtolist
func GetPackageitemsigninfodtolist() *Packageitemsigninfodtolist {
	return poolPackageitemsigninfodtolist.Get().(*Packageitemsigninfodtolist)
}

// ReleasePackageitemsigninfodtolist 释放Packageitemsigninfodtolist
func ReleasePackageitemsigninfodtolist(v *Packageitemsigninfodtolist) {
	v.Status = ""
	v.ItemQuantity = 0
	poolPackageitemsigninfodtolist.Put(v)
}
