package tmallgenie

// Dtreturnmessage 结构体
type Dtreturnmessage struct {
	// 标识
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}
