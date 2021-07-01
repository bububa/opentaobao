package tuike

// TaOfferSearchResult 结构体
type TaOfferSearchResult struct {
	// 查询总记录数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 当前条数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 错误信息
	Errors string `json:"errors,omitempty" xml:"errors,omitempty"`
	// 请求状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 数据
	DataList []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
	// 当前页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}
