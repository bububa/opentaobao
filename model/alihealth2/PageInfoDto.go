package alihealth2

// PageInfoDto 结构体
type PageInfoDto struct {
	// result
	Results []AlibabaalihealthtracecodesellerbillresultsearchResult `json:"results,omitempty" xml:"results>alibabaalihealthtracecodesellerbillresultsearch_result,omitempty"`
	// 总数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// page
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// pageSize
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
