package media

// PageQueryResult 结构体
type PageQueryResult struct {
	// resultList
	ResultList []Resultlist `json:"result_list,omitempty" xml:"result_list>resultlist,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
