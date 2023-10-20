package tmallservice

import (
	"sync"
)

// TmallFuwuHomedecorationSupplyruleCategoryworkerlistResult 结构体
type TmallFuwuHomedecorationSupplyruleCategoryworkerlistResult struct {
	// 品牌品类工人数据
	DataList []SupplyCategoryWorkerDto `json:"data_list,omitempty" xml:"data_list>supply_category_worker_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 当前页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallFuwuHomedecorationSupplyruleCategoryworkerlistResult = sync.Pool{
	New: func() any {
		return new(TmallFuwuHomedecorationSupplyruleCategoryworkerlistResult)
	},
}

// GetTmallFuwuHomedecorationSupplyruleCategoryworkerlistResult() 从对象池中获取TmallFuwuHomedecorationSupplyruleCategoryworkerlistResult
func GetTmallFuwuHomedecorationSupplyruleCategoryworkerlistResult() *TmallFuwuHomedecorationSupplyruleCategoryworkerlistResult {
	return poolTmallFuwuHomedecorationSupplyruleCategoryworkerlistResult.Get().(*TmallFuwuHomedecorationSupplyruleCategoryworkerlistResult)
}

// ReleaseTmallFuwuHomedecorationSupplyruleCategoryworkerlistResult 释放TmallFuwuHomedecorationSupplyruleCategoryworkerlistResult
func ReleaseTmallFuwuHomedecorationSupplyruleCategoryworkerlistResult(v *TmallFuwuHomedecorationSupplyruleCategoryworkerlistResult) {
	v.DataList = v.DataList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.PageIndex = 0
	v.PageSize = 0
	v.TotalCount = 0
	v.Success = false
	poolTmallFuwuHomedecorationSupplyruleCategoryworkerlistResult.Put(v)
}
