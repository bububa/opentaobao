package drugtrace

// PageInfoDto 结构体
type PageInfoDto struct {
	// 返回结果
	ResultList []CodeActiveProcessDto `json:"result_list,omitempty" xml:"result_list>code_active_process_dto,omitempty"`
	// 返回结果对象
	Result []PEntParDto `json:"result,omitempty" xml:"result>p_ent_par_dto,omitempty"`
	// 总条数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 每页几条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 第几页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}
