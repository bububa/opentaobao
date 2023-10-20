package charity

import (
	"sync"
)

// QueryCharityTimeTopApiHsfRequest 结构体
type QueryCharityTimeTopApiHsfRequest struct {
	// 第三方userKey 必传
	UserKey string `json:"user_key,omitempty" xml:"user_key,omitempty"`
	// appkey 必传
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 活动id 必传
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

var poolQueryCharityTimeTopApiHsfRequest = sync.Pool{
	New: func() any {
		return new(QueryCharityTimeTopApiHsfRequest)
	},
}

// GetQueryCharityTimeTopApiHsfRequest() 从对象池中获取QueryCharityTimeTopApiHsfRequest
func GetQueryCharityTimeTopApiHsfRequest() *QueryCharityTimeTopApiHsfRequest {
	return poolQueryCharityTimeTopApiHsfRequest.Get().(*QueryCharityTimeTopApiHsfRequest)
}

// ReleaseQueryCharityTimeTopApiHsfRequest 释放QueryCharityTimeTopApiHsfRequest
func ReleaseQueryCharityTimeTopApiHsfRequest(v *QueryCharityTimeTopApiHsfRequest) {
	v.UserKey = ""
	v.AppKey = ""
	v.ActivityId = 0
	poolQueryCharityTimeTopApiHsfRequest.Put(v)
}
