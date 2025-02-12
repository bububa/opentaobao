package tbk

import (
	"sync"
)

// LkMaterialDto 结构体
type LkMaterialDto struct {
	// 物料链接，可以为url或淘口令
	MaterialUrl string `json:"material_url,omitempty" xml:"material_url,omitempty"`
	// 优惠券id（邀约制权限，仅面向KA淘宝客）
	CouponId string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	// 是否指定券，1-指定 0-不指定 默认为0（邀约制权限，仅面向KA淘宝客）
	IsTargetCoupon int64 `json:"is_target_coupon,omitempty" xml:"is_target_coupon,omitempty"`
}

var poolLkMaterialDto = sync.Pool{
	New: func() any {
		return new(LkMaterialDto)
	},
}

// GetLkMaterialDto() 从对象池中获取LkMaterialDto
func GetLkMaterialDto() *LkMaterialDto {
	return poolLkMaterialDto.Get().(*LkMaterialDto)
}

// ReleaseLkMaterialDto 释放LkMaterialDto
func ReleaseLkMaterialDto(v *LkMaterialDto) {
	v.MaterialUrl = ""
	v.CouponId = ""
	v.IsTargetCoupon = 0
	poolLkMaterialDto.Put(v)
}
