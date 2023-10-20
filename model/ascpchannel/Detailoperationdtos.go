package ascpchannel

import (
	"sync"
)

// Detailoperationdtos 结构体
type Detailoperationdtos struct {
	// 销售市场、IPM：1000
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 库存总量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 协议ID
	TradeInvId string `json:"trade_inv_id,omitempty" xml:"trade_inv_id,omitempty"`
	// 操作子单
	DetailOrderDto *Detailorderdto `json:"detail_order_dto,omitempty" xml:"detail_order_dto,omitempty"`
	// 期货计划信息
	FuturePlanInfoDto *Futureplaninfodto `json:"future_plan_info_dto,omitempty" xml:"future_plan_info_dto,omitempty"`
	// 货品信息
	ItemDto *Itemdto `json:"item_dto,omitempty" xml:"item_dto,omitempty"`
	// 位置信息
	LocationDto *Locationdto `json:"location_dto,omitempty" xml:"location_dto,omitempty"`
	// 货主信息
	OwnerDto *Ownerdto `json:"owner_dto,omitempty" xml:"owner_dto,omitempty"`
	// 策略
	StrategyDto *Strategydto `json:"strategy_dto,omitempty" xml:"strategy_dto,omitempty"`
	// 库存类型
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 附加数据信息
	AdditionalInfoDto *Additionalinfodto `json:"additional_info_dto,omitempty" xml:"additional_info_dto,omitempty"`
}

var poolDetailoperationdtos = sync.Pool{
	New: func() any {
		return new(Detailoperationdtos)
	},
}

// GetDetailoperationdtos() 从对象池中获取Detailoperationdtos
func GetDetailoperationdtos() *Detailoperationdtos {
	return poolDetailoperationdtos.Get().(*Detailoperationdtos)
}

// ReleaseDetailoperationdtos 释放Detailoperationdtos
func ReleaseDetailoperationdtos(v *Detailoperationdtos) {
	v.ChannelCode = ""
	v.Quantity = ""
	v.TradeInvId = ""
	v.DetailOrderDto = nil
	v.FuturePlanInfoDto = nil
	v.ItemDto = nil
	v.LocationDto = nil
	v.OwnerDto = nil
	v.StrategyDto = nil
	v.InventoryType = 0
	v.AdditionalInfoDto = nil
	poolDetailoperationdtos.Put(v)
}
