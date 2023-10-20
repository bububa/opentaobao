package servicecenter

import (
	"sync"
)

// AuthorizedAccountWrapper 结构体
type AuthorizedAccountWrapper struct {
	// 商家子账号列表
	SubUsers []SubUser `json:"sub_users,omitempty" xml:"sub_users>sub_user,omitempty"`
	// 商家子账号记录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolAuthorizedAccountWrapper = sync.Pool{
	New: func() any {
		return new(AuthorizedAccountWrapper)
	},
}

// GetAuthorizedAccountWrapper() 从对象池中获取AuthorizedAccountWrapper
func GetAuthorizedAccountWrapper() *AuthorizedAccountWrapper {
	return poolAuthorizedAccountWrapper.Get().(*AuthorizedAccountWrapper)
}

// ReleaseAuthorizedAccountWrapper 释放AuthorizedAccountWrapper
func ReleaseAuthorizedAccountWrapper(v *AuthorizedAccountWrapper) {
	v.SubUsers = v.SubUsers[:0]
	v.TotalCount = 0
	poolAuthorizedAccountWrapper.Put(v)
}
