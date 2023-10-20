package drugtrace

import (
	"sync"
)

// Codeandparentlist 结构体
type Codeandparentlist struct {
	// 追溯码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 码级别
	CodeLevel string `json:"code_level,omitempty" xml:"code_level,omitempty"`
	// 父码
	ParentCode string `json:"parent_code,omitempty" xml:"parent_code,omitempty"`
}

var poolCodeandparentlist = sync.Pool{
	New: func() any {
		return new(Codeandparentlist)
	},
}

// GetCodeandparentlist() 从对象池中获取Codeandparentlist
func GetCodeandparentlist() *Codeandparentlist {
	return poolCodeandparentlist.Get().(*Codeandparentlist)
}

// ReleaseCodeandparentlist 释放Codeandparentlist
func ReleaseCodeandparentlist(v *Codeandparentlist) {
	v.Code = ""
	v.CodeLevel = ""
	v.ParentCode = ""
	poolCodeandparentlist.Put(v)
}
