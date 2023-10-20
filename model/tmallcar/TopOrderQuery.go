package tmallcar

import (
	"sync"
)

// TopOrderQuery 结构体
type TopOrderQuery struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 当前页
	PageNo int64 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// 每页个数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 订单号
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolTopOrderQuery = sync.Pool{
	New: func() any {
		return new(TopOrderQuery)
	},
}

// GetTopOrderQuery() 从对象池中获取TopOrderQuery
func GetTopOrderQuery() *TopOrderQuery {
	return poolTopOrderQuery.Get().(*TopOrderQuery)
}

// ReleaseTopOrderQuery 释放TopOrderQuery
func ReleaseTopOrderQuery(v *TopOrderQuery) {
	v.StartTime = ""
	v.EndTime = ""
	v.PageNo = 0
	v.PageSize = 0
	v.OrderId = 0
	poolTopOrderQuery.Put(v)
}
