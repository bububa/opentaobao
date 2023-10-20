package moziacl

import (
	"sync"
)

// BucPrincipalParam 结构体
type BucPrincipalParam struct {
	// 操作人所在租户
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 操作人userId
	TenantId int64 `json:"tenant_id,omitempty" xml:"tenant_id,omitempty"`
}

var poolBucPrincipalParam = sync.Pool{
	New: func() any {
		return new(BucPrincipalParam)
	},
}

// GetBucPrincipalParam() 从对象池中获取BucPrincipalParam
func GetBucPrincipalParam() *BucPrincipalParam {
	return poolBucPrincipalParam.Get().(*BucPrincipalParam)
}

// ReleaseBucPrincipalParam 释放BucPrincipalParam
func ReleaseBucPrincipalParam(v *BucPrincipalParam) {
	v.UserId = 0
	v.TenantId = 0
	poolBucPrincipalParam.Put(v)
}
