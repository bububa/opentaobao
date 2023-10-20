package happytrip

import (
	"sync"
)

// Managers 结构体
type Managers struct {
	// 联系方式
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 联系人名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolManagers = sync.Pool{
	New: func() any {
		return new(Managers)
	},
}

// GetManagers() 从对象池中获取Managers
func GetManagers() *Managers {
	return poolManagers.Get().(*Managers)
}

// ReleaseManagers 释放Managers
func ReleaseManagers(v *Managers) {
	v.Tel = ""
	v.Name = ""
	poolManagers.Put(v)
}
