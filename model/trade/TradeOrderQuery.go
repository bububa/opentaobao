package trade

import (
	"sync"
)

// TradeOrderQuery 结构体
type TradeOrderQuery struct {
	// 业务订单标识
	BizOrderIds []string `json:"biz_order_ids,omitempty" xml:"biz_order_ids>string,omitempty"`
	// 查询变动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 查询变动开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 门店标识
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 买家标识
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小（不大于20）
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolTradeOrderQuery = sync.Pool{
	New: func() any {
		return new(TradeOrderQuery)
	},
}

// GetTradeOrderQuery() 从对象池中获取TradeOrderQuery
func GetTradeOrderQuery() *TradeOrderQuery {
	return poolTradeOrderQuery.Get().(*TradeOrderQuery)
}

// ReleaseTradeOrderQuery 释放TradeOrderQuery
func ReleaseTradeOrderQuery(v *TradeOrderQuery) {
	v.BizOrderIds = v.BizOrderIds[:0]
	v.EndTime = ""
	v.StartTime = ""
	v.ShopId = ""
	v.UserId = ""
	v.PageIndex = 0
	v.PageSize = 0
	poolTradeOrderQuery.Put(v)
}
