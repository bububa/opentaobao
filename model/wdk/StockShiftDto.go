package wdk

import (
	"sync"
)

// StockShiftDto 结构体
type StockShiftDto struct {
	// itemList
	ItemList []StockShiftDetailDto `json:"item_list,omitempty" xml:"item_list>stock_shift_detail_dto,omitempty"`
	// remark
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// shiftDescribe
	ShiftDescribe string `json:"shift_describe,omitempty" xml:"shift_describe,omitempty"`
	// occurrenceDate
	OccurrenceDate string `json:"occurrence_date,omitempty" xml:"occurrence_date,omitempty"`
	// warehouseCode
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// documentNo
	DocumentNo string `json:"document_no,omitempty" xml:"document_no,omitempty"`
	// uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

var poolStockShiftDto = sync.Pool{
	New: func() any {
		return new(StockShiftDto)
	},
}

// GetStockShiftDto() 从对象池中获取StockShiftDto
func GetStockShiftDto() *StockShiftDto {
	return poolStockShiftDto.Get().(*StockShiftDto)
}

// ReleaseStockShiftDto 释放StockShiftDto
func ReleaseStockShiftDto(v *StockShiftDto) {
	v.ItemList = v.ItemList[:0]
	v.Remark = ""
	v.ShiftDescribe = ""
	v.OccurrenceDate = ""
	v.WarehouseCode = ""
	v.DocumentNo = ""
	v.Uuid = ""
	poolStockShiftDto.Put(v)
}
