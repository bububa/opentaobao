package wlbimports

import (
	"sync"
)

// HandoverPackageInfo 结构体
type HandoverPackageInfo struct {
	// 大包长(CM)，7位以内的正整数
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 大包宽(CM)，7位以内的正整数
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 大包重量（KG），7位以内的正整数
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 大包高(CM)，7位以内的正整数
	Height string `json:"height,omitempty" xml:"height,omitempty"`
}

var poolHandoverPackageInfo = sync.Pool{
	New: func() any {
		return new(HandoverPackageInfo)
	},
}

// GetHandoverPackageInfo() 从对象池中获取HandoverPackageInfo
func GetHandoverPackageInfo() *HandoverPackageInfo {
	return poolHandoverPackageInfo.Get().(*HandoverPackageInfo)
}

// ReleaseHandoverPackageInfo 释放HandoverPackageInfo
func ReleaseHandoverPackageInfo(v *HandoverPackageInfo) {
	v.Length = ""
	v.Width = ""
	v.Weight = ""
	v.Height = ""
	poolHandoverPackageInfo.Put(v)
}
