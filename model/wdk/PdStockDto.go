package wdk

import (
	"sync"
)

// PdStockDto 结构体
type PdStockDto struct {
	// itemList
	ItemList []PdStockDetailDto `json:"item_list,omitempty" xml:"item_list>pd_stock_detail_dto,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 单据号
	PdOrderCode string `json:"pd_order_code,omitempty" xml:"pd_order_code,omitempty"`
	// 店仓code，指的是盘点对象，对应一个物理店或仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 唯一识别码
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 盘点类型，可选值：1：大盘  2：周盘 ；
	PdType int64 `json:"pd_type,omitempty" xml:"pd_type,omitempty"`
}

var poolPdStockDto = sync.Pool{
	New: func() any {
		return new(PdStockDto)
	},
}

// GetPdStockDto() 从对象池中获取PdStockDto
func GetPdStockDto() *PdStockDto {
	return poolPdStockDto.Get().(*PdStockDto)
}

// ReleasePdStockDto 释放PdStockDto
func ReleasePdStockDto(v *PdStockDto) {
	v.ItemList = v.ItemList[:0]
	v.Remark = ""
	v.PdOrderCode = ""
	v.WarehouseCode = ""
	v.Uuid = ""
	v.PdType = 0
	poolPdStockDto.Put(v)
}
