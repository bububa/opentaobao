package tmallgenie

// MessageSendTarget 结构体
type MessageSendTarget struct {
	// 消息发送用户标识类型
	TargetType string `json:"target_type,omitempty" xml:"target_type,omitempty"`
	// 消息发送用户标识id
	TargetIdentity string `json:"target_identity,omitempty" xml:"target_identity,omitempty"`
}
