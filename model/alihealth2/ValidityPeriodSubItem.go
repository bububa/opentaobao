package alihealth2

import (
	"sync"
)

// ValidityPeriodSubItem 结构体
type ValidityPeriodSubItem struct {
	// 生产日期
	ProduceDate string `json:"produce_date,omitempty" xml:"produce_date,omitempty"`
	// 失效日期
	ExpireDate string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 生产批号
	ProduceNo string `json:"produce_no,omitempty" xml:"produce_no,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolValidityPeriodSubItem = sync.Pool{
	New: func() any {
		return new(ValidityPeriodSubItem)
	},
}

// GetValidityPeriodSubItem() 从对象池中获取ValidityPeriodSubItem
func GetValidityPeriodSubItem() *ValidityPeriodSubItem {
	return poolValidityPeriodSubItem.Get().(*ValidityPeriodSubItem)
}

// ReleaseValidityPeriodSubItem 释放ValidityPeriodSubItem
func ReleaseValidityPeriodSubItem(v *ValidityPeriodSubItem) {
	v.ProduceDate = ""
	v.ExpireDate = ""
	v.ProduceNo = ""
	v.Quantity = 0
	poolValidityPeriodSubItem.Put(v)
}
