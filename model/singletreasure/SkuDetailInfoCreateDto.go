package singletreasure

import (
	"sync"
)

// SkuDetailInfoCreateDto 结构体
type SkuDetailInfoCreateDto struct {
	// 猫客折上折，优惠力度，打折、减钱：单位分；打折，8折：800
	MkDiscount int64 `json:"mk_discount,omitempty" xml:"mk_discount,omitempty"`
	// sku的Id
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 优惠力度，打折、减钱：单位分；打折，8折：800
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
}

var poolSkuDetailInfoCreateDto = sync.Pool{
	New: func() any {
		return new(SkuDetailInfoCreateDto)
	},
}

// GetSkuDetailInfoCreateDto() 从对象池中获取SkuDetailInfoCreateDto
func GetSkuDetailInfoCreateDto() *SkuDetailInfoCreateDto {
	return poolSkuDetailInfoCreateDto.Get().(*SkuDetailInfoCreateDto)
}

// ReleaseSkuDetailInfoCreateDto 释放SkuDetailInfoCreateDto
func ReleaseSkuDetailInfoCreateDto(v *SkuDetailInfoCreateDto) {
	v.MkDiscount = 0
	v.SkuId = 0
	v.Discount = 0
	poolSkuDetailInfoCreateDto.Put(v)
}
