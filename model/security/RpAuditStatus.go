package security

import (
	"sync"
)

// RpAuditStatus 结构体
type RpAuditStatus struct {
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// name
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolRpAuditStatus = sync.Pool{
	New: func() any {
		return new(RpAuditStatus)
	},
}

// GetRpAuditStatus() 从对象池中获取RpAuditStatus
func GetRpAuditStatus() *RpAuditStatus {
	return poolRpAuditStatus.Get().(*RpAuditStatus)
}

// ReleaseRpAuditStatus 释放RpAuditStatus
func ReleaseRpAuditStatus(v *RpAuditStatus) {
	v.Desc = ""
	v.Name = ""
	v.Code = 0
	poolRpAuditStatus.Put(v)
}
