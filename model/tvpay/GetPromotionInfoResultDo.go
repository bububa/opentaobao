package tvpay

import (
	"sync"
)

// GetPromotionInfoResultDo 结构体
type GetPromotionInfoResultDo struct {
	// 描述
	Hint string `json:"hint,omitempty" xml:"hint,omitempty"`
	// 是否有抽奖活动
	HasPromotionEvent bool `json:"has_promotion_event,omitempty" xml:"has_promotion_event,omitempty"`
}

var poolGetPromotionInfoResultDo = sync.Pool{
	New: func() any {
		return new(GetPromotionInfoResultDo)
	},
}

// GetGetPromotionInfoResultDo() 从对象池中获取GetPromotionInfoResultDo
func GetGetPromotionInfoResultDo() *GetPromotionInfoResultDo {
	return poolGetPromotionInfoResultDo.Get().(*GetPromotionInfoResultDo)
}

// ReleaseGetPromotionInfoResultDo 释放GetPromotionInfoResultDo
func ReleaseGetPromotionInfoResultDo(v *GetPromotionInfoResultDo) {
	v.Hint = ""
	v.HasPromotionEvent = false
	poolGetPromotionInfoResultDo.Put(v)
}
