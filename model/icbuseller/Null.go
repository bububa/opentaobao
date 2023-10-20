package icbuseller

import (
	"sync"
)

// Null 结构体
type Null struct {
}

var poolNull = sync.Pool{
	New: func() any {
		return new(Null)
	},
}

// GetNull() 从对象池中获取Null
func GetNull() *Null {
	return poolNull.Get().(*Null)
}

// ReleaseNull 释放Null
func ReleaseNull(v *Null) {
	poolNull.Put(v)
}
