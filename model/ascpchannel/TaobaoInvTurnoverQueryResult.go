package ascpchannel

// TaobaoinvturnoverqueryResult 结构体
type TaobaoinvturnoverqueryResult struct {
	// 流水信息&lt;明细&gt;
	DataList []TaobaoinvturnoverqueryData `json:"data_list,omitempty" xml:"data_list>taobaoinvturnoverquery_data,omitempty"`
	// 查询页
	PageIndex string `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 单页记录数
	PageSize string `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总记录数
	TotalCount string `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页数
	CurrentPageCount string `json:"current_page_count,omitempty" xml:"current_page_count,omitempty"`
	// 总页数
	TotalPage string `json:"total_page,omitempty" xml:"total_page,omitempty"`
}
