package mos

import (
	"sync"
)

// ConfPickupGoodsReqDto 结构体
type ConfPickupGoodsReqDto struct {
	// 操作员
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 核销码
	ShortCode string `json:"short_code,omitempty" xml:"short_code,omitempty"`
	// OC交易号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 门店号
	OutBelongId string `json:"out_belong_id,omitempty" xml:"out_belong_id,omitempty"`
	// 专柜号
	OutStoreId string `json:"out_store_id,omitempty" xml:"out_store_id,omitempty"`
}

var poolConfPickupGoodsReqDto = sync.Pool{
	New: func() any {
		return new(ConfPickupGoodsReqDto)
	},
}

// GetConfPickupGoodsReqDto() 从对象池中获取ConfPickupGoodsReqDto
func GetConfPickupGoodsReqDto() *ConfPickupGoodsReqDto {
	return poolConfPickupGoodsReqDto.Get().(*ConfPickupGoodsReqDto)
}

// ReleaseConfPickupGoodsReqDto 释放ConfPickupGoodsReqDto
func ReleaseConfPickupGoodsReqDto(v *ConfPickupGoodsReqDto) {
	v.Operator = ""
	v.ShortCode = ""
	v.TradeNo = ""
	v.OutBelongId = ""
	v.OutStoreId = ""
	poolConfPickupGoodsReqDto.Put(v)
}
