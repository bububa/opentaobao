package iot

// ContentTagDto 
type ContentTagDto struct {
    // 标签id
    TagId   int64 `json:"tag_id,omitempty" xml:"tag_id,omitempty"`
    // 标签名称
    TagName   string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
}
