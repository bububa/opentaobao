package alicom

import (
	"sync"
)

// Interface 结构体
type Interface struct {
	// 接口类型
	ApiType string `json:"api_type,omitempty" xml:"api_type,omitempty"`
	// 接口地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolInterface = sync.Pool{
	New: func() any {
		return new(Interface)
	},
}

// GetInterface() 从对象池中获取Interface
func GetInterface() *Interface {
	return poolInterface.Get().(*Interface)
}

// ReleaseInterface 释放Interface
func ReleaseInterface(v *Interface) {
	v.ApiType = ""
	v.Url = ""
	poolInterface.Put(v)
}
