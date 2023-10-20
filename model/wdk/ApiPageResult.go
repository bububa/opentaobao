package wdk

import (
	"sync"
)

// ApiPageResult 结构体
type ApiPageResult struct {
	// 返回内容
	Models []CouponStatisticsResultDo `json:"models,omitempty" xml:"models>coupon_statistics_result_do,omitempty"`
	// 错误编码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 每页记录数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 是否还有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolApiPageResult = sync.Pool{
	New: func() any {
		return new(ApiPageResult)
	},
}

// GetApiPageResult() 从对象池中获取ApiPageResult
func GetApiPageResult() *ApiPageResult {
	return poolApiPageResult.Get().(*ApiPageResult)
}

// ReleaseApiPageResult 释放ApiPageResult
func ReleaseApiPageResult(v *ApiPageResult) {
	v.Models = v.Models[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = ""
	v.PageIndex = 0
	v.PageSize = 0
	v.HasNext = false
	poolApiPageResult.Put(v)
}
