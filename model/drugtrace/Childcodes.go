package drugtrace

import (
	"sync"
)

// Childcodes 结构体
type Childcodes struct {
	// 子码级别
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 子码
	PkgLevel int64 `json:"pkg_level,omitempty" xml:"pkg_level,omitempty"`
}

var poolChildcodes = sync.Pool{
	New: func() any {
		return new(Childcodes)
	},
}

// GetChildcodes() 从对象池中获取Childcodes
func GetChildcodes() *Childcodes {
	return poolChildcodes.Get().(*Childcodes)
}

// ReleaseChildcodes 释放Childcodes
func ReleaseChildcodes(v *Childcodes) {
	v.Code = ""
	v.PkgLevel = 0
	poolChildcodes.Put(v)
}
