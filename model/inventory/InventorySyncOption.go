package inventory

import (
	"sync"
)

// InventorySyncOption 结构体
type InventorySyncOption struct {
	// 业务
	BizName string `json:"biz_name,omitempty" xml:"biz_name,omitempty"`
}

var poolInventorySyncOption = sync.Pool{
	New: func() any {
		return new(InventorySyncOption)
	},
}

// GetInventorySyncOption() 从对象池中获取InventorySyncOption
func GetInventorySyncOption() *InventorySyncOption {
	return poolInventorySyncOption.Get().(*InventorySyncOption)
}

// ReleaseInventorySyncOption 释放InventorySyncOption
func ReleaseInventorySyncOption(v *InventorySyncOption) {
	v.BizName = ""
	poolInventorySyncOption.Put(v)
}
