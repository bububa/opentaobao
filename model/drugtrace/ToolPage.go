package drugtrace

// ToolPage 结构体
type ToolPage struct {
	// 返回列表
	Data []LogCodeReplaceDto `json:"data,omitempty" xml:"data>log_code_replace_dto,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页数
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
