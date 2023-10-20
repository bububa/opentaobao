package game

import (
	"sync"
)

// AppleCardCancelDto 结构体
type AppleCardCancelDto struct {
	// 单卡取消激活结果描述
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 单卡取消激活结果码
	StatusCode string `json:"status_code,omitempty" xml:"status_code,omitempty"`
	// 面值
	FacePrice string `json:"face_price,omitempty" xml:"face_price,omitempty"`
	// 卡号
	CardNo string `json:"card_no,omitempty" xml:"card_no,omitempty"`
	// 产品编码
	ZhxGoodsId string `json:"zhx_goods_id,omitempty" xml:"zhx_goods_id,omitempty"`
}

var poolAppleCardCancelDto = sync.Pool{
	New: func() any {
		return new(AppleCardCancelDto)
	},
}

// GetAppleCardCancelDto() 从对象池中获取AppleCardCancelDto
func GetAppleCardCancelDto() *AppleCardCancelDto {
	return poolAppleCardCancelDto.Get().(*AppleCardCancelDto)
}

// ReleaseAppleCardCancelDto 释放AppleCardCancelDto
func ReleaseAppleCardCancelDto(v *AppleCardCancelDto) {
	v.StatusDesc = ""
	v.StatusCode = ""
	v.FacePrice = ""
	v.CardNo = ""
	v.ZhxGoodsId = ""
	poolAppleCardCancelDto.Put(v)
}
