package tmallservice

import (
	"sync"
)

// TmallFuwuHomedecorationWorkerlevelListResult 结构体
type TmallFuwuHomedecorationWorkerlevelListResult struct {
	// 工人分层数据
	DataList []SupplyLevelWorkerDto `json:"data_list,omitempty" xml:"data_list>supply_level_worker_dto,omitempty"`
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

var poolTmallFuwuHomedecorationWorkerlevelListResult = sync.Pool{
	New: func() any {
		return new(TmallFuwuHomedecorationWorkerlevelListResult)
	},
}

// GetTmallFuwuHomedecorationWorkerlevelListResult() 从对象池中获取TmallFuwuHomedecorationWorkerlevelListResult
func GetTmallFuwuHomedecorationWorkerlevelListResult() *TmallFuwuHomedecorationWorkerlevelListResult {
	return poolTmallFuwuHomedecorationWorkerlevelListResult.Get().(*TmallFuwuHomedecorationWorkerlevelListResult)
}

// ReleaseTmallFuwuHomedecorationWorkerlevelListResult 释放TmallFuwuHomedecorationWorkerlevelListResult
func ReleaseTmallFuwuHomedecorationWorkerlevelListResult(v *TmallFuwuHomedecorationWorkerlevelListResult) {
	v.DataList = v.DataList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.PageIndex = 0
	v.PageSize = 0
	v.TotalCount = 0
	v.Success = false
	poolTmallFuwuHomedecorationWorkerlevelListResult.Put(v)
}
