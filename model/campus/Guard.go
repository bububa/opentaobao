package campus

import (
	"sync"
)

// Guard 结构体
type Guard struct {
	// 门禁点名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 门禁点id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolGuard = sync.Pool{
	New: func() any {
		return new(Guard)
	},
}

// GetGuard() 从对象池中获取Guard
func GetGuard() *Guard {
	return poolGuard.Get().(*Guard)
}

// ReleaseGuard 释放Guard
func ReleaseGuard(v *Guard) {
	v.Name = ""
	v.Id = 0
	poolGuard.Put(v)
}
