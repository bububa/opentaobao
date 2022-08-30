package user

// ReplyMessageDto 结构体
type ReplyMessageDto struct {
	// text or mix
	ContentType string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 来源消息 id
	OriginalMessageId string `json:"original_message_id,omitempty" xml:"original_message_id,omitempty"`
	// 被回复消息发送者
	ReceiverId string `json:"receiver_id,omitempty" xml:"receiver_id,omitempty"`
	// 消息内容结构
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 业务消息标记（接受回执消息）
	BizMessageTag string `json:"biz_message_tag,omitempty" xml:"biz_message_tag,omitempty"`
	// 小程序Id
	AppId string `json:"app_id,omitempty" xml:"app_id,omitempty"`
	// 毫秒时间戳
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
}
