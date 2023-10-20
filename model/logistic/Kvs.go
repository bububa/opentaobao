package logistic

import (
	"sync"
)

// Kvs 结构体
type Kvs struct {
	// value
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// key
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// index
	IndexId int64 `json:"index_id,omitempty" xml:"index_id,omitempty"`
}

var poolKvs = sync.Pool{
	New: func() any {
		return new(Kvs)
	},
}

// GetKvs() 从对象池中获取Kvs
func GetKvs() *Kvs {
	return poolKvs.Get().(*Kvs)
}

// ReleaseKvs 释放Kvs
func ReleaseKvs(v *Kvs) {
	v.Value = ""
	v.Key = ""
	v.IndexId = 0
	poolKvs.Put(v)
}
