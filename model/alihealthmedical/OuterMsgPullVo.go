package alihealthmedical

// OuterMsgPullVo 结构体
type OuterMsgPullVo struct {
	// 消息id
	RecordId string `json:"record_id,omitempty" xml:"record_id,omitempty"`
	// 会话id
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
}
