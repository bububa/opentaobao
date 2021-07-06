package openim

// EsMessage 结构体
type EsMessage struct {
	// 消息内容
	Content []RoamingMessageItem `json:"content,omitempty" xml:"content>roaming_message_item,omitempty"`
	// 消息时间，UTC时间
	Time int64 `json:"time,omitempty" xml:"time,omitempty"`
	// 消息UUID
	Uuid int64 `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 消息类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 发送方
	FromId *OpenImUser `json:"from_id,omitempty" xml:"from_id,omitempty"`
	// 接收方
	ToId *OpenImUser `json:"to_id,omitempty" xml:"to_id,omitempty"`
}
