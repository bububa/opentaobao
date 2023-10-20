package drugtrace

import (
	"sync"
)

// Parentcodeinfolist 结构体
type Parentcodeinfolist struct {
	// 父码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
	// 父码级别
	PackageLevel string `json:"package_level,omitempty" xml:"package_level,omitempty"`
}

var poolParentcodeinfolist = sync.Pool{
	New: func() any {
		return new(Parentcodeinfolist)
	},
}

// GetParentcodeinfolist() 从对象池中获取Parentcodeinfolist
func GetParentcodeinfolist() *Parentcodeinfolist {
	return poolParentcodeinfolist.Get().(*Parentcodeinfolist)
}

// ReleaseParentcodeinfolist 释放Parentcodeinfolist
func ReleaseParentcodeinfolist(v *Parentcodeinfolist) {
	v.ParentCode = ""
	v.PackageLevel = ""
	poolParentcodeinfolist.Put(v)
}
