package wdk

import (
	"sync"
)

// ItemInfoDto 结构体
type ItemInfoDto struct {
	// 加工服务
	ServiceNames []string `json:"service_names,omitempty" xml:"service_names>string,omitempty"`
	// 库存单位
	SkuStockUnit string `json:"sku_stock_unit,omitempty" xml:"sku_stock_unit,omitempty"`
	// 商品原单价
	ItemUnitPrice string `json:"item_unit_price,omitempty" xml:"item_unit_price,omitempty"`
	// 商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 实称重量
	NonstandardItemCount string `json:"nonstandard_item_count,omitempty" xml:"nonstandard_item_count,omitempty"`
	// 商品原合计金额
	TotalAmount string `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 商品应拣重量
	ExpectStockQuantity string `json:"expect_stock_quantity,omitempty" xml:"expect_stock_quantity,omitempty"`
	// 商品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 条形码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 缺货出件数
	OutOfStockItemCount int64 `json:"out_of_stock_item_count,omitempty" xml:"out_of_stock_item_count,omitempty"`
	// 商品应拣数量
	ExpectItemCount int64 `json:"expect_item_count,omitempty" xml:"expect_item_count,omitempty"`
	// 温层信息
	StorageMode int64 `json:"storage_mode,omitempty" xml:"storage_mode,omitempty"`
	// 是否标品
	StandardSku bool `json:"standard_sku,omitempty" xml:"standard_sku,omitempty"`
}

var poolItemInfoDto = sync.Pool{
	New: func() any {
		return new(ItemInfoDto)
	},
}

// GetItemInfoDto() 从对象池中获取ItemInfoDto
func GetItemInfoDto() *ItemInfoDto {
	return poolItemInfoDto.Get().(*ItemInfoDto)
}

// ReleaseItemInfoDto 释放ItemInfoDto
func ReleaseItemInfoDto(v *ItemInfoDto) {
	v.ServiceNames = v.ServiceNames[:0]
	v.SkuStockUnit = ""
	v.ItemUnitPrice = ""
	v.ItemCode = ""
	v.NonstandardItemCount = ""
	v.TotalAmount = ""
	v.ExpectStockQuantity = ""
	v.ItemName = ""
	v.Barcode = ""
	v.OutOfStockItemCount = 0
	v.ExpectItemCount = 0
	v.StorageMode = 0
	v.StandardSku = false
	poolItemInfoDto.Put(v)
}
