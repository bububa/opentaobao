package campus

import (
	"sync"
)

// UserLocationInfoQuery 结构体
type UserLocationInfoQuery struct {
	// 用户id集合
	UserIds []string `json:"user_ids,omitempty" xml:"user_ids>string,omitempty"`
	// 用户id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户身份信息
	Identity string `json:"identity,omitempty" xml:"identity,omitempty"`
	// 用户身份信息类型(1代表传入的是userId,2代表mac地址,3代表支付宝id)
	IdentityType int64 `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
	// 开始时间
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
}

var poolUserLocationInfoQuery = sync.Pool{
	New: func() any {
		return new(UserLocationInfoQuery)
	},
}

// GetUserLocationInfoQuery() 从对象池中获取UserLocationInfoQuery
func GetUserLocationInfoQuery() *UserLocationInfoQuery {
	return poolUserLocationInfoQuery.Get().(*UserLocationInfoQuery)
}

// ReleaseUserLocationInfoQuery 释放UserLocationInfoQuery
func ReleaseUserLocationInfoQuery(v *UserLocationInfoQuery) {
	v.UserIds = v.UserIds[:0]
	v.UserId = ""
	v.Identity = ""
	v.IdentityType = 0
	v.StartTime = 0
	v.EndTime = 0
	poolUserLocationInfoQuery.Put(v)
}
