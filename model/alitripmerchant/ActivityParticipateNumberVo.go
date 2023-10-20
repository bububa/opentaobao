package alitripmerchant

import (
	"sync"
)

// ActivityParticipateNumberVo 结构体
type ActivityParticipateNumberVo struct {
	// 活动id
	OfferId int64 `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
	// 剩余次数
	LeftTime int64 `json:"left_time,omitempty" xml:"left_time,omitempty"`
}

var poolActivityParticipateNumberVo = sync.Pool{
	New: func() any {
		return new(ActivityParticipateNumberVo)
	},
}

// GetActivityParticipateNumberVo() 从对象池中获取ActivityParticipateNumberVo
func GetActivityParticipateNumberVo() *ActivityParticipateNumberVo {
	return poolActivityParticipateNumberVo.Get().(*ActivityParticipateNumberVo)
}

// ReleaseActivityParticipateNumberVo 释放ActivityParticipateNumberVo
func ReleaseActivityParticipateNumberVo(v *ActivityParticipateNumberVo) {
	v.OfferId = 0
	v.LeftTime = 0
	poolActivityParticipateNumberVo.Put(v)
}
