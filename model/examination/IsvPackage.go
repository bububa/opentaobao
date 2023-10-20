package examination

import (
	"sync"
)

// IsvPackage 结构体
type IsvPackage struct {
	// 套餐可用加项id列表
	AddIsvItemIds []string `json:"add_isv_item_ids,omitempty" xml:"add_isv_item_ids>string,omitempty"`
	// 套餐可用加项包id列表
	AddIsvPackIds []string `json:"add_isv_pack_ids,omitempty" xml:"add_isv_pack_ids>string,omitempty"`
	// 套餐基础项id列表
	BasicItemIds []string `json:"basic_item_ids,omitempty" xml:"basic_item_ids>string,omitempty"`
	// 套餐代码，机构保证全局唯一
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
}

var poolIsvPackage = sync.Pool{
	New: func() any {
		return new(IsvPackage)
	},
}

// GetIsvPackage() 从对象池中获取IsvPackage
func GetIsvPackage() *IsvPackage {
	return poolIsvPackage.Get().(*IsvPackage)
}

// ReleaseIsvPackage 释放IsvPackage
func ReleaseIsvPackage(v *IsvPackage) {
	v.AddIsvItemIds = v.AddIsvItemIds[:0]
	v.AddIsvPackIds = v.AddIsvPackIds[:0]
	v.BasicItemIds = v.BasicItemIds[:0]
	v.PackageCode = ""
	poolIsvPackage.Put(v)
}
