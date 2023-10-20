package moziacl

import (
	"sync"
)

// BucUser 结构体
type BucUser struct {
	// userId
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolBucUser = sync.Pool{
	New: func() any {
		return new(BucUser)
	},
}

// GetBucUser() 从对象池中获取BucUser
func GetBucUser() *BucUser {
	return poolBucUser.Get().(*BucUser)
}

// ReleaseBucUser 释放BucUser
func ReleaseBucUser(v *BucUser) {
	v.UserId = 0
	poolBucUser.Put(v)
}
