package tbk

import (
	"sync"
)

// LkItemDto 结构体
type LkItemDto struct {
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 优惠券id（邀约制权限，仅面向KA淘宝客）
	CouponId string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	// 淘宝客外部用户标记，如自身系统账户ID；微信ID等
	ExternalId string `json:"external_id,omitempty" xml:"external_id,omitempty"`
	// 1表示商品转通用计划链接，其他值或不传表示转最优佣金率（含营销计划）链接
	Dx string `json:"dx,omitempty" xml:"dx,omitempty"`
	// 入参商品id下的skuid，传入时会透传至转链结果url中
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 是否指定券，1-指定 0-不指定 默认为0（邀约制权限，仅面向KA淘宝客）
	IsTargetCoupon int64 `json:"is_target_coupon,omitempty" xml:"is_target_coupon,omitempty"`
	// 商品库账号ID
	ManagePubId int64 `json:"manage_pub_id,omitempty" xml:"manage_pub_id,omitempty"`
}

var poolLkItemDto = sync.Pool{
	New: func() any {
		return new(LkItemDto)
	},
}

// GetLkItemDto() 从对象池中获取LkItemDto
func GetLkItemDto() *LkItemDto {
	return poolLkItemDto.Get().(*LkItemDto)
}

// ReleaseLkItemDto 释放LkItemDto
func ReleaseLkItemDto(v *LkItemDto) {
	v.ItemId = ""
	v.CouponId = ""
	v.ExternalId = ""
	v.Dx = ""
	v.SkuId = 0
	v.IsTargetCoupon = 0
	v.ManagePubId = 0
	poolLkItemDto.Put(v)
}
