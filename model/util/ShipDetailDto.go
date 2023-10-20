package util

import (
	"sync"
)

// ShipDetailDto 结构体
type ShipDetailDto struct {
	// 外部订单ID
	OuterTradeNo string `json:"outer_trade_no,omitempty" xml:"outer_trade_no,omitempty"`
	// 商品
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 0发货成功 1发货失败 10 核销成功 20 核销失败
	ShipStatus int64 `json:"ship_status,omitempty" xml:"ship_status,omitempty"`
}

var poolShipDetailDto = sync.Pool{
	New: func() any {
		return new(ShipDetailDto)
	},
}

// GetShipDetailDto() 从对象池中获取ShipDetailDto
func GetShipDetailDto() *ShipDetailDto {
	return poolShipDetailDto.Get().(*ShipDetailDto)
}

// ReleaseShipDetailDto 释放ShipDetailDto
func ReleaseShipDetailDto(v *ShipDetailDto) {
	v.OuterTradeNo = ""
	v.ItemId = 0
	v.ShipStatus = 0
	poolShipDetailDto.Put(v)
}
