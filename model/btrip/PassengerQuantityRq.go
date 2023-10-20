package btrip

import (
	"sync"
)

// PassengerQuantityRq 结构体
type PassengerQuantityRq struct {
	// 乘客类型;ADT:&#34;普通成人&#34;, CHD:&#34;儿童&#34;, STU:&#34;留学生&#34;, LABOR:&#34;劳工&#34;, MIGRANT:&#34;新移民&#34;, MARINER:&#34;海员&#34;, OLD:&#34;老人&#34;, YOUNG:&#34;青年&#34;, INFANT:&#34;婴儿&#34;, OTHER:&#34;特殊身份&#34;
	PassengerType string `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// 人员数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolPassengerQuantityRq = sync.Pool{
	New: func() any {
		return new(PassengerQuantityRq)
	},
}

// GetPassengerQuantityRq() 从对象池中获取PassengerQuantityRq
func GetPassengerQuantityRq() *PassengerQuantityRq {
	return poolPassengerQuantityRq.Get().(*PassengerQuantityRq)
}

// ReleasePassengerQuantityRq 释放PassengerQuantityRq
func ReleasePassengerQuantityRq(v *PassengerQuantityRq) {
	v.PassengerType = ""
	v.Quantity = 0
	poolPassengerQuantityRq.Put(v)
}
