package idle

import (
	"sync"
)

// TenderOrderQuery 结构体
type TenderOrderQuery struct {
	// 订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolTenderOrderQuery = sync.Pool{
	New: func() any {
		return new(TenderOrderQuery)
	},
}

// GetTenderOrderQuery() 从对象池中获取TenderOrderQuery
func GetTenderOrderQuery() *TenderOrderQuery {
	return poolTenderOrderQuery.Get().(*TenderOrderQuery)
}

// ReleaseTenderOrderQuery 释放TenderOrderQuery
func ReleaseTenderOrderQuery(v *TenderOrderQuery) {
	v.BizOrderId = 0
	poolTenderOrderQuery.Put(v)
}
