package alihealth2

import (
	"sync"
)

// PrescriptionItem 结构体
type PrescriptionItem struct {
	// 用法用量，启用用法用量规则的条件下必传
	Dosage *Dosage `json:"dosage,omitempty" xml:"dosage,omitempty"`
	// 购买数量，启用限购的规则条件下必传
	BuyAmount *BuyAmount `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 购买的药品
	Drug *Drug `json:"drug,omitempty" xml:"drug,omitempty"`
}

var poolPrescriptionItem = sync.Pool{
	New: func() any {
		return new(PrescriptionItem)
	},
}

// GetPrescriptionItem() 从对象池中获取PrescriptionItem
func GetPrescriptionItem() *PrescriptionItem {
	return poolPrescriptionItem.Get().(*PrescriptionItem)
}

// ReleasePrescriptionItem 释放PrescriptionItem
func ReleasePrescriptionItem(v *PrescriptionItem) {
	v.Dosage = nil
	v.BuyAmount = nil
	v.Drug = nil
	poolPrescriptionItem.Put(v)
}
