package alsc

import (
	"sync"
)

// EntityPrizeTokenResp 结构体
type EntityPrizeTokenResp struct {
	// 权益ID
	RightId string `json:"right_id,omitempty" xml:"right_id,omitempty"`
	// 权益实例ID
	RightInstanceId string `json:"right_instance_id,omitempty" xml:"right_instance_id,omitempty"`
}

var poolEntityPrizeTokenResp = sync.Pool{
	New: func() any {
		return new(EntityPrizeTokenResp)
	},
}

// GetEntityPrizeTokenResp() 从对象池中获取EntityPrizeTokenResp
func GetEntityPrizeTokenResp() *EntityPrizeTokenResp {
	return poolEntityPrizeTokenResp.Get().(*EntityPrizeTokenResp)
}

// ReleaseEntityPrizeTokenResp 释放EntityPrizeTokenResp
func ReleaseEntityPrizeTokenResp(v *EntityPrizeTokenResp) {
	v.RightId = ""
	v.RightInstanceId = ""
	poolEntityPrizeTokenResp.Put(v)
}
