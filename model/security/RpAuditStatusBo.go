package security

import (
	"sync"
)

// RpAuditStatusBo 结构体
type RpAuditStatusBo struct {
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// code
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
}

var poolRpAuditStatusBo = sync.Pool{
	New: func() any {
		return new(RpAuditStatusBo)
	},
}

// GetRpAuditStatusBo() 从对象池中获取RpAuditStatusBo
func GetRpAuditStatusBo() *RpAuditStatusBo {
	return poolRpAuditStatusBo.Get().(*RpAuditStatusBo)
}

// ReleaseRpAuditStatusBo 释放RpAuditStatusBo
func ReleaseRpAuditStatusBo(v *RpAuditStatusBo) {
	v.Desc = ""
	v.Name = ""
	v.Code = 0
	poolRpAuditStatusBo.Put(v)
}
