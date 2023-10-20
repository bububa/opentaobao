package qimen

import (
	"sync"
)

// SnReportRequest 结构体
type SnReportRequest struct {
	// 商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 总页数
	TotalPage int64 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
	// 当前页(从1开始)
	CurrentPage int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// 每页记录的条数
	PageSize int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 发货单信息
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenSnReportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolSnReportRequest = sync.Pool{
	New: func() any {
		return new(SnReportRequest)
	},
}

// GetSnReportRequest() 从对象池中获取SnReportRequest
func GetSnReportRequest() *SnReportRequest {
	return poolSnReportRequest.Get().(*SnReportRequest)
}

// ReleaseSnReportRequest 释放SnReportRequest
func ReleaseSnReportRequest(v *SnReportRequest) {
	v.Items = v.Items[:0]
	v.TotalPage = 0
	v.CurrentPage = 0
	v.PageSize = 0
	v.DeliveryOrder = nil
	v.ExtendProps = nil
	poolSnReportRequest.Put(v)
}
