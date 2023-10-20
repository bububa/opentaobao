package mozi

import (
	"sync"
)

// ListAccountsByAccountIdsRequest 结构体
type ListAccountsByAccountIdsRequest struct {
	// 账号ID列表
	AccountIds []int64 `json:"account_ids,omitempty" xml:"account_ids>int64,omitempty"`
	// 账号是否可用
	Available string `json:"available,omitempty" xml:"available,omitempty"`
	// 附加信息
	RequestMetaData string `json:"request_meta_data,omitempty" xml:"request_meta_data,omitempty"`
	// 租户ID
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolListAccountsByAccountIdsRequest = sync.Pool{
	New: func() any {
		return new(ListAccountsByAccountIdsRequest)
	},
}

// GetListAccountsByAccountIdsRequest() 从对象池中获取ListAccountsByAccountIdsRequest
func GetListAccountsByAccountIdsRequest() *ListAccountsByAccountIdsRequest {
	return poolListAccountsByAccountIdsRequest.Get().(*ListAccountsByAccountIdsRequest)
}

// ReleaseListAccountsByAccountIdsRequest 释放ListAccountsByAccountIdsRequest
func ReleaseListAccountsByAccountIdsRequest(v *ListAccountsByAccountIdsRequest) {
	v.AccountIds = v.AccountIds[:0]
	v.Available = ""
	v.RequestMetaData = ""
	v.TenantId = 0
	poolListAccountsByAccountIdsRequest.Put(v)
}
