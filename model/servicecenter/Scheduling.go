package servicecenter

import (
	"sync"
)

// Scheduling 结构体
type Scheduling struct {
	// 排班起始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 排班结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 商家子账号
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 服务商子账号
	SpNick string `json:"sp_nick,omitempty" xml:"sp_nick,omitempty"`
	// 排班记录状态描述
	StateDes string `json:"state_des,omitempty" xml:"state_des,omitempty"`
	// 排班记录状态，1表示生效，-1表示失效
	State int64 `json:"state,omitempty" xml:"state,omitempty"`
}

var poolScheduling = sync.Pool{
	New: func() any {
		return new(Scheduling)
	},
}

// GetScheduling() 从对象池中获取Scheduling
func GetScheduling() *Scheduling {
	return poolScheduling.Get().(*Scheduling)
}

// ReleaseScheduling 释放Scheduling
func ReleaseScheduling(v *Scheduling) {
	v.StartTime = ""
	v.EndTime = ""
	v.SellerNick = ""
	v.SpNick = ""
	v.StateDes = ""
	v.State = 0
	poolScheduling.Put(v)
}
