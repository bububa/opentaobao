package alihealthmedical

import (
	"sync"
)

// OuterMsgPullVo 结构体
type OuterMsgPullVo struct {
	// 消息id
	RecordId string `json:"record_id,omitempty" xml:"record_id,omitempty"`
	// 会话id
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
}

var poolOuterMsgPullVo = sync.Pool{
	New: func() any {
		return new(OuterMsgPullVo)
	},
}

// GetOuterMsgPullVo() 从对象池中获取OuterMsgPullVo
func GetOuterMsgPullVo() *OuterMsgPullVo {
	return poolOuterMsgPullVo.Get().(*OuterMsgPullVo)
}

// ReleaseOuterMsgPullVo 释放OuterMsgPullVo
func ReleaseOuterMsgPullVo(v *OuterMsgPullVo) {
	v.RecordId = ""
	v.SessionId = ""
	poolOuterMsgPullVo.Put(v)
}
