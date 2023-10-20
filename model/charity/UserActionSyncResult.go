package charity

import (
	"sync"
)

// UserActionSyncResult 结构体
type UserActionSyncResult struct {
	// 唯一的事件ID
	EventId string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// 本次同步获得的公益时
	CharityHours string `json:"charity_hours,omitempty" xml:"charity_hours,omitempty"`
}

var poolUserActionSyncResult = sync.Pool{
	New: func() any {
		return new(UserActionSyncResult)
	},
}

// GetUserActionSyncResult() 从对象池中获取UserActionSyncResult
func GetUserActionSyncResult() *UserActionSyncResult {
	return poolUserActionSyncResult.Get().(*UserActionSyncResult)
}

// ReleaseUserActionSyncResult 释放UserActionSyncResult
func ReleaseUserActionSyncResult(v *UserActionSyncResult) {
	v.EventId = ""
	v.CharityHours = ""
	poolUserActionSyncResult.Put(v)
}
