package alihealth2

import (
	"sync"
)

// Dosage 结构体
type Dosage struct {
	// 用法用量
	Items []DosageItem `json:"items,omitempty" xml:"items>dosage_item,omitempty"`
	// 用法
	Way string `json:"way,omitempty" xml:"way,omitempty"`
}

var poolDosage = sync.Pool{
	New: func() any {
		return new(Dosage)
	},
}

// GetDosage() 从对象池中获取Dosage
func GetDosage() *Dosage {
	return poolDosage.Get().(*Dosage)
}

// ReleaseDosage 释放Dosage
func ReleaseDosage(v *Dosage) {
	v.Items = v.Items[:0]
	v.Way = ""
	poolDosage.Put(v)
}
