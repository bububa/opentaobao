package ascpffo

import (
	"sync"
)

// ErpWarehouseInventoryDto 结构体
type ErpWarehouseInventoryDto struct {
	// 仓名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 仓编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 货品条形码
	WhcBarCode string `json:"whc_bar_code,omitempty" xml:"whc_bar_code,omitempty"`
	// 货品编码
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 货品名称
	ScItemName string `json:"sc_item_name,omitempty" xml:"sc_item_name,omitempty"`
	// 占用数量
	LockQuantity string `json:"lock_quantity,omitempty" xml:"lock_quantity,omitempty"`
	// 仓内实际库存数量(包含占用)
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 货品Id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 库存类型(1 良品，101 残品)
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
}

var poolErpWarehouseInventoryDto = sync.Pool{
	New: func() any {
		return new(ErpWarehouseInventoryDto)
	},
}

// GetErpWarehouseInventoryDto() 从对象池中获取ErpWarehouseInventoryDto
func GetErpWarehouseInventoryDto() *ErpWarehouseInventoryDto {
	return poolErpWarehouseInventoryDto.Get().(*ErpWarehouseInventoryDto)
}

// ReleaseErpWarehouseInventoryDto 释放ErpWarehouseInventoryDto
func ReleaseErpWarehouseInventoryDto(v *ErpWarehouseInventoryDto) {
	v.StoreName = ""
	v.StoreCode = ""
	v.WhcBarCode = ""
	v.ScItemCode = ""
	v.ScItemName = ""
	v.LockQuantity = ""
	v.Quantity = ""
	v.Feature = ""
	v.ScItemId = 0
	v.InventoryType = 0
	poolErpWarehouseInventoryDto.Put(v)
}
