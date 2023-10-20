package alihealthmedical

import (
	"sync"
)

// OuterMsgPullRequest 结构体
type OuterMsgPullRequest struct {
	// 外部医生id
	DoctorUuid string `json:"doctor_uuid,omitempty" xml:"doctor_uuid,omitempty"`
	// 会话id
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 内容
	Content *OuterMsgContent `json:"content,omitempty" xml:"content,omitempty"`
	// 消息内容类型  1-文本，2-图片，3-音频，31-医嘱小结
	ContentType int64 `json:"content_type,omitempty" xml:"content_type,omitempty"`
}

var poolOuterMsgPullRequest = sync.Pool{
	New: func() any {
		return new(OuterMsgPullRequest)
	},
}

// GetOuterMsgPullRequest() 从对象池中获取OuterMsgPullRequest
func GetOuterMsgPullRequest() *OuterMsgPullRequest {
	return poolOuterMsgPullRequest.Get().(*OuterMsgPullRequest)
}

// ReleaseOuterMsgPullRequest 释放OuterMsgPullRequest
func ReleaseOuterMsgPullRequest(v *OuterMsgPullRequest) {
	v.DoctorUuid = ""
	v.SessionId = ""
	v.Content = nil
	v.ContentType = 0
	poolOuterMsgPullRequest.Put(v)
}
