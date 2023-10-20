package lstlogistics

import (
	"sync"
)

// LstLogisticsInfoQuery 结构体
type LstLogisticsInfoQuery struct {
	// 主订单id
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}

var poolLstLogisticsInfoQuery = sync.Pool{
	New: func() any {
		return new(LstLogisticsInfoQuery)
	},
}

// GetLstLogisticsInfoQuery() 从对象池中获取LstLogisticsInfoQuery
func GetLstLogisticsInfoQuery() *LstLogisticsInfoQuery {
	return poolLstLogisticsInfoQuery.Get().(*LstLogisticsInfoQuery)
}

// ReleaseLstLogisticsInfoQuery 释放LstLogisticsInfoQuery
func ReleaseLstLogisticsInfoQuery(v *LstLogisticsInfoQuery) {
	v.MainOrderId = 0
	poolLstLogisticsInfoQuery.Put(v)
}
