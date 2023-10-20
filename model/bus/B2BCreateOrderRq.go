package bus

import (
	"sync"
)

// B2BCreateOrderRq 结构体
type B2BCreateOrderRq struct {
	// 乘客信息
	Passengers []PassengerVo `json:"passengers,omitempty" xml:"passengers>passenger_vo,omitempty"`
	// 取票人
	B2BFetchHolderInfo *B2BFetchHolderInfo `json:"b2_b_fetch_holder_info,omitempty" xml:"b2_b_fetch_holder_info,omitempty"`
	// 车次信息
	B2bBusLineInfo *B2BBusLineInfo `json:"b2b_bus_line_info,omitempty" xml:"b2b_bus_line_info,omitempty"`
	// 票数
	TicketCount int64 `json:"ticket_count,omitempty" xml:"ticket_count,omitempty"`
	// 总价
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
}

var poolB2BCreateOrderRq = sync.Pool{
	New: func() any {
		return new(B2BCreateOrderRq)
	},
}

// GetB2BCreateOrderRq() 从对象池中获取B2BCreateOrderRq
func GetB2BCreateOrderRq() *B2BCreateOrderRq {
	return poolB2BCreateOrderRq.Get().(*B2BCreateOrderRq)
}

// ReleaseB2BCreateOrderRq 释放B2BCreateOrderRq
func ReleaseB2BCreateOrderRq(v *B2BCreateOrderRq) {
	v.Passengers = v.Passengers[:0]
	v.B2BFetchHolderInfo = nil
	v.B2bBusLineInfo = nil
	v.TicketCount = 0
	v.TotalPrice = 0
	poolB2BCreateOrderRq.Put(v)
}
