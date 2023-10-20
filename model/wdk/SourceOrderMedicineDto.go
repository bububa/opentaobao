package wdk

import (
	"sync"
)

// SourceOrderMedicineDto 结构体
type SourceOrderMedicineDto struct {
	// 药品明细list
	MedicineItemDOS []MedicineItemDo `json:"medicine_item_d_o_s,omitempty" xml:"medicine_item_d_o_s>medicine_item_do,omitempty"`
	// 盒马履约单号
	SourceOrderCode string `json:"source_order_code,omitempty" xml:"source_order_code,omitempty"`
	// 外部单号(比如饿了么)
	OriginalOrderId string `json:"original_order_id,omitempty" xml:"original_order_id,omitempty"`
	// 盒马app订单号
	OutMainOrderId string `json:"out_main_order_id,omitempty" xml:"out_main_order_id,omitempty"`
}

var poolSourceOrderMedicineDto = sync.Pool{
	New: func() any {
		return new(SourceOrderMedicineDto)
	},
}

// GetSourceOrderMedicineDto() 从对象池中获取SourceOrderMedicineDto
func GetSourceOrderMedicineDto() *SourceOrderMedicineDto {
	return poolSourceOrderMedicineDto.Get().(*SourceOrderMedicineDto)
}

// ReleaseSourceOrderMedicineDto 释放SourceOrderMedicineDto
func ReleaseSourceOrderMedicineDto(v *SourceOrderMedicineDto) {
	v.MedicineItemDOS = v.MedicineItemDOS[:0]
	v.SourceOrderCode = ""
	v.OriginalOrderId = ""
	v.OutMainOrderId = ""
	poolSourceOrderMedicineDto.Put(v)
}
