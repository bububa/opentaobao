package qimen

import (
	"sync"
)

// Packages 结构体
type Packages struct {
	// 包裹详情
	PackageValue *Package `json:"packageValue,omitempty" xml:"packageValue,omitempty"`
}

var poolPackages = sync.Pool{
	New: func() any {
		return new(Packages)
	},
}

// GetPackages() 从对象池中获取Packages
func GetPackages() *Packages {
	return poolPackages.Get().(*Packages)
}

// ReleasePackages 释放Packages
func ReleasePackages(v *Packages) {
	v.PackageValue = nil
	poolPackages.Put(v)
}
