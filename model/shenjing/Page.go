package shenjing

// Page 结构体
type Page struct {
	// 活动列表
	Items []AlibabashenjingcoreactivitygetappshowlistT `json:"items,omitempty" xml:"items>alibabashenjingcoreactivitygetappshowlist_t,omitempty"`
	// 分页总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 一页行数
	Limit int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// 分页总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 一页行数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}
