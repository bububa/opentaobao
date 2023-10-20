package tbk

import (
	"sync"
)

// TopApiRefundRptOption 结构体
type TopApiRefundRptOption struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// pagesize
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 1-维权发起时间，2-订单结算时间（正向订单），3-维权完成时间，4-订单创建时间，5-订单更新时间
	SearchType int64 `json:"search_type,omitempty" xml:"search_type,omitempty"`
	// 1 表示2方，2表示3方，0表示不限
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	// pagenumber
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 1代表渠道关系id，2代表会员关系id
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

var poolTopApiRefundRptOption = sync.Pool{
	New: func() any {
		return new(TopApiRefundRptOption)
	},
}

// GetTopApiRefundRptOption() 从对象池中获取TopApiRefundRptOption
func GetTopApiRefundRptOption() *TopApiRefundRptOption {
	return poolTopApiRefundRptOption.Get().(*TopApiRefundRptOption)
}

// ReleaseTopApiRefundRptOption 释放TopApiRefundRptOption
func ReleaseTopApiRefundRptOption(v *TopApiRefundRptOption) {
	v.StartTime = ""
	v.PageSize = 0
	v.SearchType = 0
	v.RefundType = 0
	v.PageNo = 0
	v.BizType = 0
	poolTopApiRefundRptOption.Put(v)
}
