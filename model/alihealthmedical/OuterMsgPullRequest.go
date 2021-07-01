package alihealthmedical

// OuterMsgPullRequest 结构体
type OuterMsgPullRequest struct {
	// 内容
	Content *OuterMsgContent `json:"content,omitempty" xml:"content,omitempty"`
	// 消息内容类型  1-文本，2-图片，3-音频，31-医嘱小结
	ContentType int64 `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 会话id
	SessionId string `json:"session_id,omitempty" xml:"session_id,omitempty"`
	// 外部医生id
	DoctorUuid string `json:"doctor_uuid,omitempty" xml:"doctor_uuid,omitempty"`
}
