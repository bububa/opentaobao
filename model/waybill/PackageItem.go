package waybill

import (
	"sync"
)

// PackageItem 结构体
type PackageItem struct {
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 商品数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolPackageItem = sync.Pool{
	New: func() any {
		return new(PackageItem)
	},
}

// GetPackageItem() 从对象池中获取PackageItem
func GetPackageItem() *PackageItem {
	return poolPackageItem.Get().(*PackageItem)
}

// ReleasePackageItem 释放PackageItem
func ReleasePackageItem(v *PackageItem) {
	v.ItemName = ""
	v.Count = 0
	poolPackageItem.Put(v)
}
