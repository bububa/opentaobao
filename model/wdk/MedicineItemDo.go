package wdk

import (
	"sync"
)

// MedicineItemDo 结构体
type MedicineItemDo struct {
	// sku名称
	SkuTitle string `json:"sku_title,omitempty" xml:"sku_title,omitempty"`
	// sku编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
}

var poolMedicineItemDo = sync.Pool{
	New: func() any {
		return new(MedicineItemDo)
	},
}

// GetMedicineItemDo() 从对象池中获取MedicineItemDo
func GetMedicineItemDo() *MedicineItemDo {
	return poolMedicineItemDo.Get().(*MedicineItemDo)
}

// ReleaseMedicineItemDo 释放MedicineItemDo
func ReleaseMedicineItemDo(v *MedicineItemDo) {
	v.SkuTitle = ""
	v.SkuCode = ""
	v.Count = 0
	poolMedicineItemDo.Put(v)
}
