package bus

import (
	"sync"
)

// TopBusNumberUpdateRq 结构体
type TopBusNumberUpdateRq struct {
	// 车次列表
	Numbers []BusNumberDto `json:"numbers,omitempty" xml:"numbers>bus_number_dto,omitempty"`
}

var poolTopBusNumberUpdateRq = sync.Pool{
	New: func() any {
		return new(TopBusNumberUpdateRq)
	},
}

// GetTopBusNumberUpdateRq() 从对象池中获取TopBusNumberUpdateRq
func GetTopBusNumberUpdateRq() *TopBusNumberUpdateRq {
	return poolTopBusNumberUpdateRq.Get().(*TopBusNumberUpdateRq)
}

// ReleaseTopBusNumberUpdateRq 释放TopBusNumberUpdateRq
func ReleaseTopBusNumberUpdateRq(v *TopBusNumberUpdateRq) {
	v.Numbers = v.Numbers[:0]
	poolTopBusNumberUpdateRq.Put(v)
}
