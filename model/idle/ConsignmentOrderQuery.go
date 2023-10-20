package idle

import (
	"sync"
)

// ConsignmentOrderQuery 结构体
type ConsignmentOrderQuery struct {
	// 订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolConsignmentOrderQuery = sync.Pool{
	New: func() any {
		return new(ConsignmentOrderQuery)
	},
}

// GetConsignmentOrderQuery() 从对象池中获取ConsignmentOrderQuery
func GetConsignmentOrderQuery() *ConsignmentOrderQuery {
	return poolConsignmentOrderQuery.Get().(*ConsignmentOrderQuery)
}

// ReleaseConsignmentOrderQuery 释放ConsignmentOrderQuery
func ReleaseConsignmentOrderQuery(v *ConsignmentOrderQuery) {
	v.BizOrderId = 0
	poolConsignmentOrderQuery.Put(v)
}
