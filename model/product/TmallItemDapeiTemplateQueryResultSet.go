package product

// TmallItemDapeiTemplateQueryResultSet 结构体
type TmallItemDapeiTemplateQueryResultSet struct {
	// firstResult
	Results []DapeiDo `json:"results,omitempty" xml:"results>dapei_do,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// totalResults
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// totalPage
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// pageIndex
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// error
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
}
