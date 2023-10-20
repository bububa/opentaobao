package alihealth2

import (
	"sync"
)

// Diag 结构体
type Diag struct {
	// 诊断码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 诊断名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

var poolDiag = sync.Pool{
	New: func() any {
		return new(Diag)
	},
}

// GetDiag() 从对象池中获取Diag
func GetDiag() *Diag {
	return poolDiag.Get().(*Diag)
}

// ReleaseDiag 释放Diag
func ReleaseDiag(v *Diag) {
	v.Code = ""
	v.Name = ""
	poolDiag.Put(v)
}
