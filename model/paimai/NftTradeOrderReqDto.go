package paimai

import (
	"sync"
)

// NftTradeOrderReqDto 结构体
type NftTradeOrderReqDto struct {
	// 订单ID多个用逗号分开
	OrderIds string `json:"order_ids,omitempty" xml:"order_ids,omitempty"`
}

var poolNftTradeOrderReqDto = sync.Pool{
	New: func() any {
		return new(NftTradeOrderReqDto)
	},
}

// GetNftTradeOrderReqDto() 从对象池中获取NftTradeOrderReqDto
func GetNftTradeOrderReqDto() *NftTradeOrderReqDto {
	return poolNftTradeOrderReqDto.Get().(*NftTradeOrderReqDto)
}

// ReleaseNftTradeOrderReqDto 释放NftTradeOrderReqDto
func ReleaseNftTradeOrderReqDto(v *NftTradeOrderReqDto) {
	v.OrderIds = ""
	poolNftTradeOrderReqDto.Put(v)
}
