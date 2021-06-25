package openim

// TribeTextMessage 
type TribeTextMessage struct {

    // 发送方userid。必须为本app已导入的账号
    FromId   string `json:"from_id,omitempty"`

    // 消息
    Message   string `json:"message,omitempty"`

    // 消息时间。UTC时间，精确到秒。时间范围必须在当前时间30天内
    Time   int64 `json:"time,omitempty"`

}
