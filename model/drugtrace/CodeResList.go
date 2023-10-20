package drugtrace

import (
	"sync"
)

// CodeResList 结构体
type CodeResList struct {
	// 码前缀
	CodePrefix string `json:"code_prefix,omitempty" xml:"code_prefix,omitempty"`
	// 资源码
	ResCode string `json:"res_code,omitempty" xml:"res_code,omitempty"`
	// 层级
	CodeLevel string `json:"code_level,omitempty" xml:"code_level,omitempty"`
	// 包装比例
	PkgRatio string `json:"pkg_ratio,omitempty" xml:"pkg_ratio,omitempty"`
}

var poolCodeResList = sync.Pool{
	New: func() any {
		return new(CodeResList)
	},
}

// GetCodeResList() 从对象池中获取CodeResList
func GetCodeResList() *CodeResList {
	return poolCodeResList.Get().(*CodeResList)
}

// ReleaseCodeResList 释放CodeResList
func ReleaseCodeResList(v *CodeResList) {
	v.CodePrefix = ""
	v.ResCode = ""
	v.CodeLevel = ""
	v.PkgRatio = ""
	poolCodeResList.Put(v)
}
