package xhotelitem

import (
	"sync"
)

// PromoCode 结构体
type PromoCode struct {
	// 营销活动code
	ActivityCode string `json:"activity_code,omitempty" xml:"activity_code,omitempty"`
}

var poolPromoCode = sync.Pool{
	New: func() any {
		return new(PromoCode)
	},
}

// GetPromoCode() 从对象池中获取PromoCode
func GetPromoCode() *PromoCode {
	return poolPromoCode.Get().(*PromoCode)
}

// ReleasePromoCode 释放PromoCode
func ReleasePromoCode(v *PromoCode) {
	v.ActivityCode = ""
	poolPromoCode.Put(v)
}
