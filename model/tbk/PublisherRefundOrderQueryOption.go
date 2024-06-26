package tbk

import (
	"sync"
)

// PublisherRefundOrderQueryOption 结构体
type PublisherRefundOrderQueryOption struct {
	// 位点，除第一页之外，都需要传递；前端原样返回。
	PositionIndex string `json:"position_index,omitempty" xml:"position_index,omitempty"`
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
	JumpType int64 `json:"jump_type,omitempty" xml:"jump_type,omitempty"`
	// 1全部订单
	OrderScene int64 `json:"order_scene,omitempty" xml:"order_scene,omitempty"`
	// 单页订单量，1-100，上限100条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 1-维权发起时间，2-订单结算时间（正向订单），3-维权完成时间，4-订单创建时间，5-订单更新时间 （注意：入参startTime小于2022年12月20日0点时，query_type只能为1、2、3；入参startTime大于等于2022年12月20日0点时，query_type可以为1、2、3、4、5)
	QueryType int64 `json:"query_type,omitempty" xml:"query_type,omitempty"`
	// 查询页数
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 1 表示2方，2表示3方，4表示不限
	MemberType int64 `json:"member_type,omitempty" xml:"member_type,omitempty"`
}

var poolPublisherRefundOrderQueryOption = sync.Pool{
	New: func() any {
		return new(PublisherRefundOrderQueryOption)
	},
}

// GetPublisherRefundOrderQueryOption() 从对象池中获取PublisherRefundOrderQueryOption
func GetPublisherRefundOrderQueryOption() *PublisherRefundOrderQueryOption {
	return poolPublisherRefundOrderQueryOption.Get().(*PublisherRefundOrderQueryOption)
}

// ReleasePublisherRefundOrderQueryOption 释放PublisherRefundOrderQueryOption
func ReleasePublisherRefundOrderQueryOption(v *PublisherRefundOrderQueryOption) {
	v.PositionIndex = ""
	v.StartTime = ""
	v.EndTime = ""
	v.JumpType = 0
	v.OrderScene = 0
	v.PageSize = 0
	v.QueryType = 0
	v.PageNo = 0
	v.MemberType = 0
	poolPublisherRefundOrderQueryOption.Put(v)
}
