package security

import (
	"sync"
)

// RpAuditValueBo 结构体
type RpAuditValueBo struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolRpAuditValueBo = sync.Pool{
	New: func() any {
		return new(RpAuditValueBo)
	},
}

// GetRpAuditValueBo() 从对象池中获取RpAuditValueBo
func GetRpAuditValueBo() *RpAuditValueBo {
	return poolRpAuditValueBo.Get().(*RpAuditValueBo)
}

// ReleaseRpAuditValueBo 释放RpAuditValueBo
func ReleaseRpAuditValueBo(v *RpAuditValueBo) {
	v.Desc = ""
	v.Name = ""
	v.Code = 0
	poolRpAuditValueBo.Put(v)
}
