package product

import (
	"sync"
)

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

var poolTmallItemDapeiTemplateQueryResultSet = sync.Pool{
	New: func() any {
		return new(TmallItemDapeiTemplateQueryResultSet)
	},
}

// GetTmallItemDapeiTemplateQueryResultSet() 从对象池中获取TmallItemDapeiTemplateQueryResultSet
func GetTmallItemDapeiTemplateQueryResultSet() *TmallItemDapeiTemplateQueryResultSet {
	return poolTmallItemDapeiTemplateQueryResultSet.Get().(*TmallItemDapeiTemplateQueryResultSet)
}

// ReleaseTmallItemDapeiTemplateQueryResultSet 释放TmallItemDapeiTemplateQueryResultSet
func ReleaseTmallItemDapeiTemplateQueryResultSet(v *TmallItemDapeiTemplateQueryResultSet) {
	v.Results = v.Results[:0]
	v.ErrorMsg = ""
	v.ErrorCode = ""
	v.TotalResults = 0
	v.TotalPage = 0
	v.PageIndex = 0
	v.Error = false
	poolTmallItemDapeiTemplateQueryResultSet.Put(v)
}
