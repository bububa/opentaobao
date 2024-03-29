package wdk

import (
	"sync"
)

// ErpBillDto 结构体
type ErpBillDto struct {
	// orderCode
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// status
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// createDate
	CreateDate string `json:"create_date,omitempty" xml:"create_date,omitempty"`
	// warehouseCode
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// type
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolErpBillDto = sync.Pool{
	New: func() any {
		return new(ErpBillDto)
	},
}

// GetErpBillDto() 从对象池中获取ErpBillDto
func GetErpBillDto() *ErpBillDto {
	return poolErpBillDto.Get().(*ErpBillDto)
}

// ReleaseErpBillDto 释放ErpBillDto
func ReleaseErpBillDto(v *ErpBillDto) {
	v.OrderCode = ""
	v.Status = ""
	v.CreateDate = ""
	v.WarehouseCode = ""
	v.Type = 0
	poolErpBillDto.Put(v)
}
