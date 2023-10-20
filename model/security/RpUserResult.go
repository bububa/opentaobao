package security

import (
	"sync"
)

// RpUserResult 结构体
type RpUserResult struct {
	// users
	Users []string `json:"users,omitempty" xml:"users>string,omitempty"`
}

var poolRpUserResult = sync.Pool{
	New: func() any {
		return new(RpUserResult)
	},
}

// GetRpUserResult() 从对象池中获取RpUserResult
func GetRpUserResult() *RpUserResult {
	return poolRpUserResult.Get().(*RpUserResult)
}

// ReleaseRpUserResult 释放RpUserResult
func ReleaseRpUserResult(v *RpUserResult) {
	v.Users = v.Users[:0]
	poolRpUserResult.Put(v)
}
