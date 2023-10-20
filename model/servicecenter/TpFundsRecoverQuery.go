package servicecenter

import (
	"sync"
)

// TpFundsRecoverQuery 结构体
type TpFundsRecoverQuery struct {
	// 订单ID
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolTpFundsRecoverQuery = sync.Pool{
	New: func() any {
		return new(TpFundsRecoverQuery)
	},
}

// GetTpFundsRecoverQuery() 从对象池中获取TpFundsRecoverQuery
func GetTpFundsRecoverQuery() *TpFundsRecoverQuery {
	return poolTpFundsRecoverQuery.Get().(*TpFundsRecoverQuery)
}

// ReleaseTpFundsRecoverQuery 释放TpFundsRecoverQuery
func ReleaseTpFundsRecoverQuery(v *TpFundsRecoverQuery) {
	v.BizOrderId = 0
	poolTpFundsRecoverQuery.Put(v)
}
