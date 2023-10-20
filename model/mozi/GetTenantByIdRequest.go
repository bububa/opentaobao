package mozi

import (
	"sync"
)

// GetTenantByIdRequest 结构体
type GetTenantByIdRequest struct {
	// 租户id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolGetTenantByIdRequest = sync.Pool{
	New: func() any {
		return new(GetTenantByIdRequest)
	},
}

// GetGetTenantByIdRequest() 从对象池中获取GetTenantByIdRequest
func GetGetTenantByIdRequest() *GetTenantByIdRequest {
	return poolGetTenantByIdRequest.Get().(*GetTenantByIdRequest)
}

// ReleaseGetTenantByIdRequest 释放GetTenantByIdRequest
func ReleaseGetTenantByIdRequest(v *GetTenantByIdRequest) {
	v.TenantId = 0
	poolGetTenantByIdRequest.Put(v)
}
