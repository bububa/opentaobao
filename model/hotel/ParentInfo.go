package hotel

import (
	"sync"
)

// ParentInfo 结构体
type ParentInfo struct {
	// 脱敏后的用户名字
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 脱敏后的userId
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolParentInfo = sync.Pool{
	New: func() any {
		return new(ParentInfo)
	},
}

// GetParentInfo() 从对象池中获取ParentInfo
func GetParentInfo() *ParentInfo {
	return poolParentInfo.Get().(*ParentInfo)
}

// ReleaseParentInfo 释放ParentInfo
func ReleaseParentInfo(v *ParentInfo) {
	v.UserNick = ""
	v.UserId = 0
	poolParentInfo.Put(v)
}
