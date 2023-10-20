package alihouse

import (
	"sync"
)

// CouponOrderQuery 结构体
type CouponOrderQuery struct {
	// 履约单id
	TradeOrderId int64 `json:"trade_order_id,omitempty" xml:"trade_order_id,omitempty"`
}

var poolCouponOrderQuery = sync.Pool{
	New: func() any {
		return new(CouponOrderQuery)
	},
}

// GetCouponOrderQuery() 从对象池中获取CouponOrderQuery
func GetCouponOrderQuery() *CouponOrderQuery {
	return poolCouponOrderQuery.Get().(*CouponOrderQuery)
}

// ReleaseCouponOrderQuery 释放CouponOrderQuery
func ReleaseCouponOrderQuery(v *CouponOrderQuery) {
	v.TradeOrderId = 0
	poolCouponOrderQuery.Put(v)
}
