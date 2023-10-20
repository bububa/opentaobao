package alitripbp

import (
	"sync"
)

// ExamineOuterUserRequest 结构体
type ExamineOuterUserRequest struct {
	// 活动标识
	Activity string `json:"activity,omitempty" xml:"activity,omitempty"`
	// 渠道标识
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 用户标识
	OuterUserId string `json:"outer_user_id,omitempty" xml:"outer_user_id,omitempty"`
	// 平台标识
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
}

var poolExamineOuterUserRequest = sync.Pool{
	New: func() any {
		return new(ExamineOuterUserRequest)
	},
}

// GetExamineOuterUserRequest() 从对象池中获取ExamineOuterUserRequest
func GetExamineOuterUserRequest() *ExamineOuterUserRequest {
	return poolExamineOuterUserRequest.Get().(*ExamineOuterUserRequest)
}

// ReleaseExamineOuterUserRequest 释放ExamineOuterUserRequest
func ReleaseExamineOuterUserRequest(v *ExamineOuterUserRequest) {
	v.Activity = ""
	v.Channel = ""
	v.OuterUserId = ""
	v.Platform = ""
	poolExamineOuterUserRequest.Put(v)
}
