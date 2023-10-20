package train

import (
	"sync"
)

// QueryOrderRs 结构体
type QueryOrderRs struct {
	// 票详情
	TicketInfos []TicketInfoDto `json:"ticket_infos,omitempty" xml:"ticket_infos>ticket_info_dto,omitempty"`
	// 子单详情
	TapSubOrders []TapSubOrderVo `json:"tap_sub_orders,omitempty" xml:"tap_sub_orders>tap_sub_order_vo,omitempty"`
	// 主单详情
	TapOrder *TapOrderVo `json:"tap_order,omitempty" xml:"tap_order,omitempty"`
}

var poolQueryOrderRs = sync.Pool{
	New: func() any {
		return new(QueryOrderRs)
	},
}

// GetQueryOrderRs() 从对象池中获取QueryOrderRs
func GetQueryOrderRs() *QueryOrderRs {
	return poolQueryOrderRs.Get().(*QueryOrderRs)
}

// ReleaseQueryOrderRs 释放QueryOrderRs
func ReleaseQueryOrderRs(v *QueryOrderRs) {
	v.TicketInfos = v.TicketInfos[:0]
	v.TapSubOrders = v.TapSubOrders[:0]
	v.TapOrder = nil
	poolQueryOrderRs.Put(v)
}
