package logistic

import (
	"sync"
)

// ChainStore 结构体
type ChainStore struct {
	// 门店code
	ChainstoreCode string `json:"chainstore_code,omitempty" xml:"chainstore_code,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
}

var poolChainStore = sync.Pool{
	New: func() any {
		return new(ChainStore)
	},
}

// GetChainStore() 从对象池中获取ChainStore
func GetChainStore() *ChainStore {
	return poolChainStore.Get().(*ChainStore)
}

// ReleaseChainStore 释放ChainStore
func ReleaseChainStore(v *ChainStore) {
	v.ChainstoreCode = ""
	v.Longitude = ""
	v.Latitude = ""
	poolChainStore.Put(v)
}
