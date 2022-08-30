package dmp

// Pager 结构体
type Pager struct {
	// 当前页数
	IntCurrentPage int64 `json:"int_current_page,omitempty" xml:"int_current_page,omitempty"`
	// 分页大小
	IntPageSize int64 `json:"int_page_size,omitempty" xml:"int_page_size,omitempty"`
	// 记录总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页数
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
