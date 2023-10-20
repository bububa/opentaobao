package wdk

import (
	"sync"
)

// StockChangeDetailDto 结构体
type StockChangeDetailDto struct {
	// quantity
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// itemCode
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// batchCode
	BatchCode string `json:"batch_code,omitempty" xml:"batch_code,omitempty"`
	// reason
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// bizOrderCode
	BizOrderCode string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
	// cabinetCode
	CabinetCode string `json:"cabinet_code,omitempty" xml:"cabinet_code,omitempty"`
}

var poolStockChangeDetailDto = sync.Pool{
	New: func() any {
		return new(StockChangeDetailDto)
	},
}

// GetStockChangeDetailDto() 从对象池中获取StockChangeDetailDto
func GetStockChangeDetailDto() *StockChangeDetailDto {
	return poolStockChangeDetailDto.Get().(*StockChangeDetailDto)
}

// ReleaseStockChangeDetailDto 释放StockChangeDetailDto
func ReleaseStockChangeDetailDto(v *StockChangeDetailDto) {
	v.Quantity = ""
	v.ItemCode = ""
	v.BatchCode = ""
	v.Reason = ""
	v.BizOrderCode = ""
	v.CabinetCode = ""
	poolStockChangeDetailDto.Put(v)
}
