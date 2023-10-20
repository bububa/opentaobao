package security

import (
	"sync"
)

// RpAuditType 结构体
type RpAuditType struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolRpAuditType = sync.Pool{
	New: func() any {
		return new(RpAuditType)
	},
}

// GetRpAuditType() 从对象池中获取RpAuditType
func GetRpAuditType() *RpAuditType {
	return poolRpAuditType.Get().(*RpAuditType)
}

// ReleaseRpAuditType 释放RpAuditType
func ReleaseRpAuditType(v *RpAuditType) {
	v.Desc = ""
	v.Name = ""
	v.Code = 0
	poolRpAuditType.Put(v)
}
