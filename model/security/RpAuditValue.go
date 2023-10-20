package security

import (
	"sync"
)

// RpAuditValue 结构体
type RpAuditValue struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolRpAuditValue = sync.Pool{
	New: func() any {
		return new(RpAuditValue)
	},
}

// GetRpAuditValue() 从对象池中获取RpAuditValue
func GetRpAuditValue() *RpAuditValue {
	return poolRpAuditValue.Get().(*RpAuditValue)
}

// ReleaseRpAuditValue 释放RpAuditValue
func ReleaseRpAuditValue(v *RpAuditValue) {
	v.Desc = ""
	v.Name = ""
	v.Code = 0
	poolRpAuditValue.Put(v)
}
