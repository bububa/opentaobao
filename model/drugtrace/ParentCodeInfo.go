package drugtrace

import (
	"sync"
)

// ParentCodeInfo 结构体
type ParentCodeInfo struct {
	// 父码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
	// 码级别
	PackageLevel int64 `json:"package_level,omitempty" xml:"package_level,omitempty"`
}

var poolParentCodeInfo = sync.Pool{
	New: func() any {
		return new(ParentCodeInfo)
	},
}

// GetParentCodeInfo() 从对象池中获取ParentCodeInfo
func GetParentCodeInfo() *ParentCodeInfo {
	return poolParentCodeInfo.Get().(*ParentCodeInfo)
}

// ReleaseParentCodeInfo 释放ParentCodeInfo
func ReleaseParentCodeInfo(v *ParentCodeInfo) {
	v.ParentCode = ""
	v.PackageLevel = 0
	poolParentCodeInfo.Put(v)
}
