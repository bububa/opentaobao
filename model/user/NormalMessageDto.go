package user

// NormalMessageDTO 
type NormalMessageDTO struct {
    // 内容
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    // 时间戳
    CreateTime   int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
    // 接收者ID
    ReceiverId   string `json:"receiver_id,omitempty" xml:"receiver_id,omitempty"`
    // 类型
    ContentType   string `json:"content_type,omitempty" xml:"content_type,omitempty"`
}
