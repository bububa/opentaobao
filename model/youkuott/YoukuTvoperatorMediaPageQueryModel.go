package youkuott

// YoukutvoperatormediapagequeryModel 结构体
type YoukutvoperatormediapagequeryModel struct {
	// 数据列表
	DataList []YoukutvoperatormediapagequeryData `json:"data_list,omitempty" xml:"data_list>youkutvoperatormediapagequery_data,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 页号
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
