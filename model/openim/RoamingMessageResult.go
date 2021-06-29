package openim

// RoamingMessageResult 
type RoamingMessageResult struct {
    // 下次迭代key
    NextKey   string `json:"next_key,omitempty" xml:"next_key,omitempty"`
    // 消息列表
    Messages   []RoamingMessage `json:"messages,omitempty" xml:"messages>roaming_message,omitempty"`
}
