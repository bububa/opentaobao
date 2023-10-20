package tblogistics

import (
	"sync"
)

// StoreInfo 结构体
type StoreInfo struct {
	// 仓库真实名字
	RealName string `json:"real_name,omitempty" xml:"real_name,omitempty"`
	// XXX果园
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 仓库服务代码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 详细地址
	Address string `json:"address,omitempty" xml:"address,omitempty"`
}

var poolStoreInfo = sync.Pool{
	New: func() any {
		return new(StoreInfo)
	},
}

// GetStoreInfo() 从对象池中获取StoreInfo
func GetStoreInfo() *StoreInfo {
	return poolStoreInfo.Get().(*StoreInfo)
}

// ReleaseStoreInfo 释放StoreInfo
func ReleaseStoreInfo(v *StoreInfo) {
	v.RealName = ""
	v.Name = ""
	v.ServiceCode = ""
	v.Address = ""
	poolStoreInfo.Put(v)
}
