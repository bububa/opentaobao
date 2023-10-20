package hotel

import (
	"sync"
)

// Headers 结构体
type Headers struct {
	// empty
	Empty bool `json:"empty,omitempty" xml:"empty,omitempty"`
}

var poolHeaders = sync.Pool{
	New: func() any {
		return new(Headers)
	},
}

// GetHeaders() 从对象池中获取Headers
func GetHeaders() *Headers {
	return poolHeaders.Get().(*Headers)
}

// ReleaseHeaders 释放Headers
func ReleaseHeaders(v *Headers) {
	v.Empty = false
	poolHeaders.Put(v)
}
