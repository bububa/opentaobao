package drugtrace

import (
	"sync"
)

// ResCode 结构体
type ResCode struct {
	// 资源码值
	Value string `json:"value,omitempty" xml:"value,omitempty"`
	// 码级别
	CodeLevel string `json:"code_level,omitempty" xml:"code_level,omitempty"`
	// 码头
	CodeVersion string `json:"code_version,omitempty" xml:"code_version,omitempty"`
	// 包装比例
	PkgRatio string `json:"pkg_ratio,omitempty" xml:"pkg_ratio,omitempty"`
}

var poolResCode = sync.Pool{
	New: func() any {
		return new(ResCode)
	},
}

// GetResCode() 从对象池中获取ResCode
func GetResCode() *ResCode {
	return poolResCode.Get().(*ResCode)
}

// ReleaseResCode 释放ResCode
func ReleaseResCode(v *ResCode) {
	v.Value = ""
	v.CodeLevel = ""
	v.CodeVersion = ""
	v.PkgRatio = ""
	poolResCode.Put(v)
}
