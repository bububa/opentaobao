package ascpchannel

import (
	"sync"
)

// DataMap 结构体
type DataMap struct {
	// 幂等原因
	IdempotentReason string `json:"idempotent_reason,omitempty" xml:"idempotent_reason,omitempty"`
}

var poolDataMap = sync.Pool{
	New: func() any {
		return new(DataMap)
	},
}

// GetDataMap() 从对象池中获取DataMap
func GetDataMap() *DataMap {
	return poolDataMap.Get().(*DataMap)
}

// ReleaseDataMap 释放DataMap
func ReleaseDataMap(v *DataMap) {
	v.IdempotentReason = ""
	poolDataMap.Put(v)
}
