package ihome

// ContentStatus 
type ContentStatus struct {
    // 拒绝原因
    RefuseReason   string `json:"refuse_reason,omitempty" xml:"refuse_reason,omitempty"`
    // 审核状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 内容ID
    ContentId   int64 `json:"content_id,omitempty" xml:"content_id,omitempty"`
    // 场景ID
    SampleId   int64 `json:"sample_id,omitempty" xml:"sample_id,omitempty"`
    // 标题
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 图片
    Img   string `json:"img,omitempty" xml:"img,omitempty"`
}
