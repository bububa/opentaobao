package hotel

import (
	"sync"
)

// NameValuePair 结构体
type NameValuePair struct {
	// 档次code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 档次名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolNameValuePair = sync.Pool{
	New: func() any {
		return new(NameValuePair)
	},
}

// GetNameValuePair() 从对象池中获取NameValuePair
func GetNameValuePair() *NameValuePair {
	return poolNameValuePair.Get().(*NameValuePair)
}

// ReleaseNameValuePair 释放NameValuePair
func ReleaseNameValuePair(v *NameValuePair) {
	v.Code = ""
	v.Name = ""
	poolNameValuePair.Put(v)
}
