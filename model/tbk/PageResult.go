package tbk

// PageResult 结构体
type PageResult struct {
	// 数据结果
	Results []TaobaotbkdgcpaactivitydetailResults `json:"results,omitempty" xml:"results>taobaotbkdgcpaactivitydetail_results,omitempty"`
	// 数据批次号(时间戳)
	Runtime string `json:"runtime,omitempty" xml:"runtime,omitempty"`
	// 上一页页码
	Prepage int64 `json:"pre_page,omitempty" xml:"pre_page,omitempty"`
	// 下一页页码
	Nextpage int64 `json:"next_page,omitempty" xml:"next_page,omitempty"`
	// 当前页码
	Pageno int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 总共页数
	Totalpages int64 `json:"total_pages,omitempty" xml:"total_pages,omitempty"`
	// 每页条数
	Pagesize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总条数
	Totalcount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否有下一页
	Hasnext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// 是否有下一页
	Haspre bool `json:"has_pre,omitempty" xml:"has_pre,omitempty"`
}
