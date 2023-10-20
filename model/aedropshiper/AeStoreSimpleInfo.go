package aedropshiper

import (
	"sync"
)

// AeStoreSimpleInfo 结构体
type AeStoreSimpleInfo struct {
	// Store name
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// Store address
	StoreUrl string `json:"store_url,omitempty" xml:"store_url,omitempty"`
	// Store ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolAeStoreSimpleInfo = sync.Pool{
	New: func() any {
		return new(AeStoreSimpleInfo)
	},
}

// GetAeStoreSimpleInfo() 从对象池中获取AeStoreSimpleInfo
func GetAeStoreSimpleInfo() *AeStoreSimpleInfo {
	return poolAeStoreSimpleInfo.Get().(*AeStoreSimpleInfo)
}

// ReleaseAeStoreSimpleInfo 释放AeStoreSimpleInfo
func ReleaseAeStoreSimpleInfo(v *AeStoreSimpleInfo) {
	v.StoreName = ""
	v.StoreUrl = ""
	v.StoreId = 0
	poolAeStoreSimpleInfo.Put(v)
}
