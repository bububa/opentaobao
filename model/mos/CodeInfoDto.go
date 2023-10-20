package mos

import (
	"sync"
)

// CodeInfoDto 结构体
type CodeInfoDto struct {
	// 商品信息
	GoodsList []CodeGoodsDto `json:"goods_list,omitempty" xml:"goods_list>code_goods_dto,omitempty"`
	// 包裹信息
	PackageCode string `json:"package_code,omitempty" xml:"package_code,omitempty"`
	// 寄件信息
	SendInfo *DeliveryCustomDto `json:"send_info,omitempty" xml:"send_info,omitempty"`
}

var poolCodeInfoDto = sync.Pool{
	New: func() any {
		return new(CodeInfoDto)
	},
}

// GetCodeInfoDto() 从对象池中获取CodeInfoDto
func GetCodeInfoDto() *CodeInfoDto {
	return poolCodeInfoDto.Get().(*CodeInfoDto)
}

// ReleaseCodeInfoDto 释放CodeInfoDto
func ReleaseCodeInfoDto(v *CodeInfoDto) {
	v.GoodsList = v.GoodsList[:0]
	v.PackageCode = ""
	v.SendInfo = nil
	poolCodeInfoDto.Put(v)
}
