package tmallgenie

// Note 
type Note struct {
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // memo_ID
    MemoId   int64 `json:"memo_id,omitempty" xml:"memo_id,omitempty"`
    // uuid
    Uuid   string `json:"uuid,omitempty" xml:"uuid,omitempty"`
    // memo状态
    Status   string `json:"status,omitempty" xml:"status,omitempty"`
    // 记事本内容
    Content   string `json:"content,omitempty" xml:"content,omitempty"`
    // 记事本主题
    Topic   string `json:"topic,omitempty" xml:"topic,omitempty"`
}
