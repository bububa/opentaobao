package paimai

import (
	"sync"
)

// NftTradeOrderDto 结构体
type NftTradeOrderDto struct {
	// 订单类型
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolNftTradeOrderDto = sync.Pool{
	New: func() any {
		return new(NftTradeOrderDto)
	},
}

// GetNftTradeOrderDto() 从对象池中获取NftTradeOrderDto
func GetNftTradeOrderDto() *NftTradeOrderDto {
	return poolNftTradeOrderDto.Get().(*NftTradeOrderDto)
}

// ReleaseNftTradeOrderDto 释放NftTradeOrderDto
func ReleaseNftTradeOrderDto(v *NftTradeOrderDto) {
	v.OrderType = ""
	v.OrderId = ""
	poolNftTradeOrderDto.Put(v)
}
