package tmallservice

import (
	"sync"
)

// PageResult 结构体
type PageResult struct {
	// 分页数据
	ResultData []SscClaimInfoDto `json:"result_data,omitempty" xml:"result_data>ssc_claim_info_dto,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页数
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPageResult = sync.Pool{
	New: func() any {
		return new(PageResult)
	},
}

// GetPageResult() 从对象池中获取PageResult
func GetPageResult() *PageResult {
	return poolPageResult.Get().(*PageResult)
}

// ReleasePageResult 释放PageResult
func ReleasePageResult(v *PageResult) {
	v.ResultData = v.ResultData[:0]
	v.MsgInfo = ""
	v.MsgCode = ""
	v.TotalCount = 0
	v.PageIndex = 0
	v.PageSize = 0
	v.Success = false
	poolPageResult.Put(v)
}
