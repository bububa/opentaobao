package openim

// TribeMessageResult 结构体
type TribeMessageResult struct {
	// 消息列表
	Messages []TribeMessage `json:"messages,omitempty" xml:"messages>tribe_message,omitempty"`
	// 迭代key
	NextKey string `json:"next_key,omitempty" xml:"next_key,omitempty"`
}
