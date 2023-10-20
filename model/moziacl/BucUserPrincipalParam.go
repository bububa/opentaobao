package moziacl

import (
	"sync"
)

// BucUserPrincipalParam 结构体
type BucUserPrincipalParam struct {
	// 租户名称，不需要填
	TenantName string `json:"tenant_name,omitempty" xml:"tenant_name,omitempty"`
	// 操作人userId
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 操作人所在租户Id
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolBucUserPrincipalParam = sync.Pool{
	New: func() any {
		return new(BucUserPrincipalParam)
	},
}

// GetBucUserPrincipalParam() 从对象池中获取BucUserPrincipalParam
func GetBucUserPrincipalParam() *BucUserPrincipalParam {
	return poolBucUserPrincipalParam.Get().(*BucUserPrincipalParam)
}

// ReleaseBucUserPrincipalParam 释放BucUserPrincipalParam
func ReleaseBucUserPrincipalParam(v *BucUserPrincipalParam) {
	v.TenantName = ""
	v.UserId = 0
	v.TenantId = 0
	poolBucUserPrincipalParam.Put(v)
}
