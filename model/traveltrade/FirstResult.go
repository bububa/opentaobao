package traveltrade

import (
	"sync"
)

// FirstResult 结构体
type FirstResult struct {
	// 预约订单ID
	BookOrderId int64 `json:"book_order_id,omitempty" xml:"book_order_id,omitempty"`
	// TC子订单号
	SubTcOrderId int64 `json:"sub_tc_order_id,omitempty" xml:"sub_tc_order_id,omitempty"`
	// 预约ID(主键)
	BookInfoId int64 `json:"book_info_id,omitempty" xml:"book_info_id,omitempty"`
	// TC主订单号
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
}

var poolFirstResult = sync.Pool{
	New: func() any {
		return new(FirstResult)
	},
}

// GetFirstResult() 从对象池中获取FirstResult
func GetFirstResult() *FirstResult {
	return poolFirstResult.Get().(*FirstResult)
}

// ReleaseFirstResult 释放FirstResult
func ReleaseFirstResult(v *FirstResult) {
	v.BookOrderId = 0
	v.SubTcOrderId = 0
	v.BookInfoId = 0
	v.TcOrderId = 0
	poolFirstResult.Put(v)
}
