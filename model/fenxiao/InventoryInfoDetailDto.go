package fenxiao

import (
	"sync"
)

// InventoryInfoDetailDto 结构体
type InventoryInfoDetailDto struct {
	// subList
	SubList []InventorySubDetailDto `json:"sub_list,omitempty" xml:"sub_list>inventory_sub_detail_dto,omitempty"`
	// 后端商品code
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// 仓库code
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 占用库存数
	OccupyQuantity int64 `json:"occupy_quantity,omitempty" xml:"occupy_quantity,omitempty"`
	// 仓库物理库存数
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 预扣库存数
	ReserveQuantity int64 `json:"reserve_quantity,omitempty" xml:"reserve_quantity,omitempty"`
	// 后端商品id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// distType
	InvStoreType int64 `json:"inv_store_type,omitempty" xml:"inv_store_type,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 1前端商品 2供应链货品
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
}

var poolInventoryInfoDetailDto = sync.Pool{
	New: func() any {
		return new(InventoryInfoDetailDto)
	},
}

// GetInventoryInfoDetailDto() 从对象池中获取InventoryInfoDetailDto
func GetInventoryInfoDetailDto() *InventoryInfoDetailDto {
	return poolInventoryInfoDetailDto.Get().(*InventoryInfoDetailDto)
}

// ReleaseInventoryInfoDetailDto 释放InventoryInfoDetailDto
func ReleaseInventoryInfoDetailDto(v *InventoryInfoDetailDto) {
	v.SubList = v.SubList[:0]
	v.ScItemCode = ""
	v.StoreCode = ""
	v.OccupyQuantity = 0
	v.Quantity = 0
	v.ReserveQuantity = 0
	v.ScItemId = 0
	v.InvStoreType = 0
	v.SkuId = 0
	v.ItemType = 0
	poolInventoryInfoDetailDto.Put(v)
}
