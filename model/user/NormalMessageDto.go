package user

// NormalMessageDto 
type NormalMessageDto struct {

    // 内容
    Content   string `json:"content,omitempty"`

    // 时间戳
    CreateTime   int64 `json:"create_time,omitempty"`

    // 接收者ID
    ReceiverId   string `json:"receiver_id,omitempty"`

    // 类型
    ContentType   string `json:"content_type,omitempty"`

}
