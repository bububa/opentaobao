package qimen

import (
	"sync"
)

// StoreKeeper 结构体
type StoreKeeper struct {
	// 传真
	Fax string `json:"fax,omitempty" xml:"fax,omitempty"`
	// 固定电话
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 邮编
	ZipCode string `json:"zip_code,omitempty" xml:"zip_code,omitempty"`
	// 移动电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

var poolStoreKeeper = sync.Pool{
	New: func() any {
		return new(StoreKeeper)
	},
}

// GetStoreKeeper() 从对象池中获取StoreKeeper
func GetStoreKeeper() *StoreKeeper {
	return poolStoreKeeper.Get().(*StoreKeeper)
}

// ReleaseStoreKeeper 释放StoreKeeper
func ReleaseStoreKeeper(v *StoreKeeper) {
	v.Fax = ""
	v.Tel = ""
	v.Name = ""
	v.ZipCode = ""
	v.Mobile = ""
	poolStoreKeeper.Put(v)
}
