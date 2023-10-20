package tmallservice

import (
	"sync"
)

// JoinedStore 结构体
type JoinedStore struct {
	// 网点编码
	ServiceStoreCode string `json:"service_store_code,omitempty" xml:"service_store_code,omitempty"`
}

var poolJoinedStore = sync.Pool{
	New: func() any {
		return new(JoinedStore)
	},
}

// GetJoinedStore() 从对象池中获取JoinedStore
func GetJoinedStore() *JoinedStore {
	return poolJoinedStore.Get().(*JoinedStore)
}

// ReleaseJoinedStore 释放JoinedStore
func ReleaseJoinedStore(v *JoinedStore) {
	v.ServiceStoreCode = ""
	poolJoinedStore.Put(v)
}
