package mozi

import (
	"sync"
)

// Tenant 结构体
type Tenant struct {
	// 租户完整code,格式：命名空间+$+code
	TenantFullCode string `json:"tenant_full_code,omitempty" xml:"tenant_full_code,omitempty"`
	// 租户名称
	TenantName string `json:"tenant_name,omitempty" xml:"tenant_name,omitempty"`
	// 租户描述
	TenantDescription string `json:"tenant_description,omitempty" xml:"tenant_description,omitempty"`
	// 创建者
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改者
	Modifier string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 租户id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolTenant = sync.Pool{
	New: func() any {
		return new(Tenant)
	},
}

// GetTenant() 从对象池中获取Tenant
func GetTenant() *Tenant {
	return poolTenant.Get().(*Tenant)
}

// ReleaseTenant 释放Tenant
func ReleaseTenant(v *Tenant) {
	v.TenantFullCode = ""
	v.TenantName = ""
	v.TenantDescription = ""
	v.Creator = ""
	v.Modifier = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.TenantId = 0
	poolTenant.Put(v)
}
