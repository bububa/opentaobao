package alihealth2

import (
	"sync"
)

// DosageItem 结构体
type DosageItem struct {
	// 使用量单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 用量类型（SINGLEDAILYDOSE(&#34;单日剂量&#34;),     SINGLEDOSE(&#34;单次剂量&#34;),     DRUGCOURSE(&#34;药品疗程&#34;),     FREQUENCYADMINISTRATION(&#34;给药频次&#34;)）
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 使用量
	Value float64 `json:"value,omitempty" xml:"value,omitempty"`
}

var poolDosageItem = sync.Pool{
	New: func() any {
		return new(DosageItem)
	},
}

// GetDosageItem() 从对象池中获取DosageItem
func GetDosageItem() *DosageItem {
	return poolDosageItem.Get().(*DosageItem)
}

// ReleaseDosageItem 释放DosageItem
func ReleaseDosageItem(v *DosageItem) {
	v.Unit = ""
	v.Type = ""
	v.Value = 0
	poolDosageItem.Put(v)
}
