package user

// MassMessageDto 
type MassMessageDto struct {

    // 创建时间
    CreateTime   int64 `json:"create_time,omitempty"`

    // 内容
    Content   string `json:"content,omitempty"`

    // 内容类型
    ContentType   string `json:"content_type,omitempty"`

    // 任务名称
    TaskName   string `json:"task_name,omitempty"`

}
