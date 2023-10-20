package drugtrace

// Page 结构体
type Page struct {
	// 返回列表
	ResultList []PentParDto `json:"result_list,omitempty" xml:"result_list>pent_par_dto,omitempty"`
	// 总数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 当前页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
