package mos

import (
	"sync"
)

// AlibabaMjMoscarnivalReceiveencryptData 结构体
type AlibabaMjMoscarnivalReceiveencryptData struct {
	// 权益列表
	RightsList []RightsList `json:"rights_list,omitempty" xml:"rights_list>rights_list,omitempty"`
	// 奖品类型
	PrizeType int64 `json:"prize_type,omitempty" xml:"prize_type,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 是否有奖品
	HasPrize bool `json:"has_prize,omitempty" xml:"has_prize,omitempty"`
	// 是否新会员
	IsNewUser bool `json:"is_new_user,omitempty" xml:"is_new_user,omitempty"`
}

var poolAlibabaMjMoscarnivalReceiveencryptData = sync.Pool{
	New: func() any {
		return new(AlibabaMjMoscarnivalReceiveencryptData)
	},
}

// GetAlibabaMjMoscarnivalReceiveencryptData() 从对象池中获取AlibabaMjMoscarnivalReceiveencryptData
func GetAlibabaMjMoscarnivalReceiveencryptData() *AlibabaMjMoscarnivalReceiveencryptData {
	return poolAlibabaMjMoscarnivalReceiveencryptData.Get().(*AlibabaMjMoscarnivalReceiveencryptData)
}

// ReleaseAlibabaMjMoscarnivalReceiveencryptData 释放AlibabaMjMoscarnivalReceiveencryptData
func ReleaseAlibabaMjMoscarnivalReceiveencryptData(v *AlibabaMjMoscarnivalReceiveencryptData) {
	v.RightsList = v.RightsList[:0]
	v.PrizeType = 0
	v.Success = false
	v.HasPrize = false
	v.IsNewUser = false
	poolAlibabaMjMoscarnivalReceiveencryptData.Put(v)
}
