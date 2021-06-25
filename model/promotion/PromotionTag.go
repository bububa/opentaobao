package promotion

// PromotionTag 
type PromotionTag struct {

    // 标签ID
    TagId   int64 `json:"tag_id,omitempty"`

    // 标签名称
    TagName   string `json:"tag_name,omitempty"`

    // 标签描述
    TagDesc   string `json:"tag_desc,omitempty"`

    // 标签开始时间
    StartTime   string `json:"start_time,omitempty"`

    // 标签结束时间
    EndTime   string `json:"end_time,omitempty"`

}
