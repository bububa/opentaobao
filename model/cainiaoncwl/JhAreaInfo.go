package cainiaoncwl

import (
	"sync"
)

// JhAreaInfo 结构体
type JhAreaInfo struct {
	// 可以指定某个省的集货单
	Provice string `json:"provice,omitempty" xml:"provice,omitempty"`
	// 可以指定某个省市的集货单
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 可以指定省市区的集货单
	City string `json:"city,omitempty" xml:"city,omitempty"`
}

var poolJhAreaInfo = sync.Pool{
	New: func() any {
		return new(JhAreaInfo)
	},
}

// GetJhAreaInfo() 从对象池中获取JhAreaInfo
func GetJhAreaInfo() *JhAreaInfo {
	return poolJhAreaInfo.Get().(*JhAreaInfo)
}

// ReleaseJhAreaInfo 释放JhAreaInfo
func ReleaseJhAreaInfo(v *JhAreaInfo) {
	v.Provice = ""
	v.District = ""
	v.City = ""
	poolJhAreaInfo.Put(v)
}
