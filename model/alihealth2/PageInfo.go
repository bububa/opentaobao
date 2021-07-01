package alihealth2

// PageInfo 结构体
type PageInfo struct {
	// totalNum
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// result
	Results []AlibabaAlihealthTracecodesellerChannelSearchResult `json:"results,omitempty" xml:"results>alibaba_alihealth_tracecodeseller_channel_search_result,omitempty"`
	// page
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// pageSize
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
