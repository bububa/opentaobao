package fenxiao

import (
	"sync"
)

// Store 结构体
type Store struct {
	// 联系电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 商家的仓库编码，不允许重复
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 仓库的物理地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
	// 仓库类型
	StoreType string `json:"store_type,omitempty" xml:"store_type,omitempty"`
	// 邮编
	PostCode string `json:"post_code,omitempty" xml:"post_code,omitempty"`
	// 联系人
	Contact string `json:"contact,omitempty" xml:"contact,omitempty"`
	// 商家的仓库名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 仓库对应的淘宝区域
	AddressAreaName string `json:"address_area_name,omitempty" xml:"address_area_name,omitempty"`
	// 仓库简称
	AliasName string `json:"alias_name,omitempty" xml:"alias_name,omitempty"`
}

var poolStore = sync.Pool{
	New: func() any {
		return new(Store)
	},
}

// GetStore() 从对象池中获取Store
func GetStore() *Store {
	return poolStore.Get().(*Store)
}

// ReleaseStore 释放Store
func ReleaseStore(v *Store) {
	v.Phone = ""
	v.StoreCode = ""
	v.Address = ""
	v.StoreType = ""
	v.PostCode = ""
	v.Contact = ""
	v.StoreName = ""
	v.AddressAreaName = ""
	v.AliasName = ""
	poolStore.Put(v)
}
