package ascpffo

import (
	"sync"
)

// ErpOnWayInventoryDto 结构体
type ErpOnWayInventoryDto struct {
	// 出库仓名称
	OutboundStoreName string `json:"outbound_store_name,omitempty" xml:"outbound_store_name,omitempty"`
	// 出库仓编码
	OutboundStoreCode string `json:"outbound_store_code,omitempty" xml:"outbound_store_code,omitempty"`
	// 入库仓名称
	InboundStoreName string `json:"inbound_store_name,omitempty" xml:"inbound_store_name,omitempty"`
	// 入库仓编码
	InboundStoreCode string `json:"inbound_store_code,omitempty" xml:"inbound_store_code,omitempty"`
	// 在途数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 货品条形码
	WhcBarCode string `json:"whc_bar_code,omitempty" xml:"whc_bar_code,omitempty"`
	// 货品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 货品名称
	ScItemName string `json:"sc_item_name,omitempty" xml:"sc_item_name,omitempty"`
	// 库存类型(1 采购在途，2 调拨在途，3 销售在途，4 销退在途)
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 货品Id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}

var poolErpOnWayInventoryDto = sync.Pool{
	New: func() any {
		return new(ErpOnWayInventoryDto)
	},
}

// GetErpOnWayInventoryDto() 从对象池中获取ErpOnWayInventoryDto
func GetErpOnWayInventoryDto() *ErpOnWayInventoryDto {
	return poolErpOnWayInventoryDto.Get().(*ErpOnWayInventoryDto)
}

// ReleaseErpOnWayInventoryDto 释放ErpOnWayInventoryDto
func ReleaseErpOnWayInventoryDto(v *ErpOnWayInventoryDto) {
	v.OutboundStoreName = ""
	v.OutboundStoreCode = ""
	v.InboundStoreName = ""
	v.InboundStoreCode = ""
	v.Quantity = ""
	v.Feature = ""
	v.WhcBarCode = ""
	v.ScItemCode = ""
	v.ScItemName = ""
	v.InventoryType = 0
	v.ScItemId = 0
	poolErpOnWayInventoryDto.Put(v)
}
