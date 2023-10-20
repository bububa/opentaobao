package exchange

import (
	"sync"
)

// RefundMessageResult 结构体
type RefundMessageResult struct {
	// 留言记录
	Results []RefundMessage `json:"results,omitempty" xml:"results>refund_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 当前页面的留言条数
	PageResults int64 `json:"page_results,omitempty" xml:"page_results,omitempty"`
	// 所有留言记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

var poolRefundMessageResult = sync.Pool{
	New: func() any {
		return new(RefundMessageResult)
	},
}

// GetRefundMessageResult() 从对象池中获取RefundMessageResult
func GetRefundMessageResult() *RefundMessageResult {
	return poolRefundMessageResult.Get().(*RefundMessageResult)
}

// ReleaseRefundMessageResult 释放RefundMessageResult
func ReleaseRefundMessageResult(v *RefundMessageResult) {
	v.Results = v.Results[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.PageResults = 0
	v.TotalResults = 0
	v.HasNext = false
	poolRefundMessageResult.Put(v)
}
