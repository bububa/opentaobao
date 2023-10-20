package lstlogistics

import (
	"sync"
)

// LstLogisticsTraceQuery 结构体
type LstLogisticsTraceQuery struct {
	// 物流编号
	LogisticsId string `json:"logistics_id,omitempty" xml:"logistics_id,omitempty"`
	// 主订单号
	MainOrderId int64 `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
}

var poolLstLogisticsTraceQuery = sync.Pool{
	New: func() any {
		return new(LstLogisticsTraceQuery)
	},
}

// GetLstLogisticsTraceQuery() 从对象池中获取LstLogisticsTraceQuery
func GetLstLogisticsTraceQuery() *LstLogisticsTraceQuery {
	return poolLstLogisticsTraceQuery.Get().(*LstLogisticsTraceQuery)
}

// ReleaseLstLogisticsTraceQuery 释放LstLogisticsTraceQuery
func ReleaseLstLogisticsTraceQuery(v *LstLogisticsTraceQuery) {
	v.LogisticsId = ""
	v.MainOrderId = 0
	poolLstLogisticsTraceQuery.Put(v)
}
