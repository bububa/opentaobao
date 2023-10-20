package bus

import (
	"sync"
)

// TopItemChangeNotifyRq 结构体
type TopItemChangeNotifyRq struct {
	// 出发城市
	FromCity string `json:"from_city,omitempty" xml:"from_city,omitempty"`
	// 目的地城市
	ToCity string `json:"to_city,omitempty" xml:"to_city,omitempty"`
	// 发车日期，格式：yyyy-MM-dd
	FromDate string `json:"from_date,omitempty" xml:"from_date,omitempty"`
}

var poolTopItemChangeNotifyRq = sync.Pool{
	New: func() any {
		return new(TopItemChangeNotifyRq)
	},
}

// GetTopItemChangeNotifyRq() 从对象池中获取TopItemChangeNotifyRq
func GetTopItemChangeNotifyRq() *TopItemChangeNotifyRq {
	return poolTopItemChangeNotifyRq.Get().(*TopItemChangeNotifyRq)
}

// ReleaseTopItemChangeNotifyRq 释放TopItemChangeNotifyRq
func ReleaseTopItemChangeNotifyRq(v *TopItemChangeNotifyRq) {
	v.FromCity = ""
	v.ToCity = ""
	v.FromDate = ""
	poolTopItemChangeNotifyRq.Put(v)
}
