package maitix

// Page 结构体
type Page struct {
	// 数据对象
	DataArrList []ProjectDto `json:"data_arr_list,omitempty" xml:"data_arr_list>project_dto,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总页数
	PageCount int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总项目数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
