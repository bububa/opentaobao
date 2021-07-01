package ieagency

// BasePageDo 结构体
type BasePageDo struct {
	// 第几页
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 分页大小
	Pageindex int64 `json:"pageindex,omitempty" xml:"pageindex,omitempty"`
	// 总纪录数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
