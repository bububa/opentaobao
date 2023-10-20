package qimen

import (
	"sync"
)

// ExpressInfo 结构体
type ExpressInfo struct {
	// 奇门仓储字段
	ExpressCode string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
	// 奇门仓储字段
	ExpressName string `json:"expressName,omitempty" xml:"expressName,omitempty"`
	// 奇门仓储字段
	BrandCode string `json:"brandCode,omitempty" xml:"brandCode,omitempty"`
	// 奇门仓储字段
	BrandName string `json:"brandName,omitempty" xml:"brandName,omitempty"`
	// 奇门仓储字段
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}

var poolExpressInfo = sync.Pool{
	New: func() any {
		return new(ExpressInfo)
	},
}

// GetExpressInfo() 从对象池中获取ExpressInfo
func GetExpressInfo() *ExpressInfo {
	return poolExpressInfo.Get().(*ExpressInfo)
}

// ReleaseExpressInfo 释放ExpressInfo
func ReleaseExpressInfo(v *ExpressInfo) {
	v.ExpressCode = ""
	v.ExpressName = ""
	v.BrandCode = ""
	v.BrandName = ""
	v.Status = ""
	poolExpressInfo.Put(v)
}
