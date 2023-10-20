package auction

import (
	"sync"
)

// Result4Top 结构体
type Result4Top struct {
	// 最新出价列表
	Results []LatestBids `json:"results,omitempty" xml:"results>latest_bids,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误说明
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回的最新出价条数
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 数据接收成功为true，失败false
	Value bool `json:"value,omitempty" xml:"value,omitempty"`
}

var poolResult4Top = sync.Pool{
	New: func() any {
		return new(Result4Top)
	},
}

// GetResult4Top() 从对象池中获取Result4Top
func GetResult4Top() *Result4Top {
	return poolResult4Top.Get().(*Result4Top)
}

// ReleaseResult4Top 释放Result4Top
func ReleaseResult4Top(v *Result4Top) {
	v.Results = v.Results[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalItem = 0
	v.Success = false
	v.Value = false
	poolResult4Top.Put(v)
}
