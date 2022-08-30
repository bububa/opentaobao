package promotion

// ActivityReadTopQuery 结构体
type ActivityReadTopQuery struct {
	// 筛选状态列表，EFFECTIVE为生效，OFFLINE为下线
	StatusList []string `json:"status_list,omitempty" xml:"status_list>string,omitempty"`
	// 发放开始时间右值
	StartTimeEnd string `json:"start_time_end,omitempty" xml:"start_time_end,omitempty"`
	// 业务来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 发放开始时间左值
	StartTimeBegin string `json:"start_time_begin,omitempty" xml:"start_time_begin,omitempty"`
	// 发放结束时间左值
	EndTimeBegin string `json:"end_time_begin,omitempty" xml:"end_time_begin,omitempty"`
	// 发放结束时间右值
	EndTimeEnd string `json:"end_time_end,omitempty" xml:"end_time_end,omitempty"`
	// 每页记录数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
}
