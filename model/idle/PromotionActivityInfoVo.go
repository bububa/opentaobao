package idle

import (
	"sync"
)

// PromotionActivityInfoVo 结构体
type PromotionActivityInfoVo struct {
	// 当前用户是否参与了营销活动
	Joined bool `json:"joined,omitempty" xml:"joined,omitempty"`
}

var poolPromotionActivityInfoVo = sync.Pool{
	New: func() any {
		return new(PromotionActivityInfoVo)
	},
}

// GetPromotionActivityInfoVo() 从对象池中获取PromotionActivityInfoVo
func GetPromotionActivityInfoVo() *PromotionActivityInfoVo {
	return poolPromotionActivityInfoVo.Get().(*PromotionActivityInfoVo)
}

// ReleasePromotionActivityInfoVo 释放PromotionActivityInfoVo
func ReleasePromotionActivityInfoVo(v *PromotionActivityInfoVo) {
	v.Joined = false
	poolPromotionActivityInfoVo.Put(v)
}
