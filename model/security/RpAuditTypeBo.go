package security

import (
	"sync"
)

// RpAuditTypeBo 结构体
type RpAuditTypeBo struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolRpAuditTypeBo = sync.Pool{
	New: func() any {
		return new(RpAuditTypeBo)
	},
}

// GetRpAuditTypeBo() 从对象池中获取RpAuditTypeBo
func GetRpAuditTypeBo() *RpAuditTypeBo {
	return poolRpAuditTypeBo.Get().(*RpAuditTypeBo)
}

// ReleaseRpAuditTypeBo 释放RpAuditTypeBo
func ReleaseRpAuditTypeBo(v *RpAuditTypeBo) {
	v.Desc = ""
	v.Name = ""
	v.Code = 0
	poolRpAuditTypeBo.Put(v)
}
