package examination

// StoreSpecialTagsResponse 结构体
type StoreSpecialTagsResponse struct {
	// 标签code
	TagCode string `json:"tag_code,omitempty" xml:"tag_code,omitempty"`
	// 标签name
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 优先级
	Priority int64 `json:"priority,omitempty" xml:"priority,omitempty"`
}
