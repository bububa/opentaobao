package fenxiao

import (
	"sync"
)

// UpdateCnskuOption 结构体
type UpdateCnskuOption struct {
	// 是否同步到wms, 为空时默认下发
	SyncWms bool `json:"sync_wms,omitempty" xml:"sync_wms,omitempty"`
}

var poolUpdateCnskuOption = sync.Pool{
	New: func() any {
		return new(UpdateCnskuOption)
	},
}

// GetUpdateCnskuOption() 从对象池中获取UpdateCnskuOption
func GetUpdateCnskuOption() *UpdateCnskuOption {
	return poolUpdateCnskuOption.Get().(*UpdateCnskuOption)
}

// ReleaseUpdateCnskuOption 释放UpdateCnskuOption
func ReleaseUpdateCnskuOption(v *UpdateCnskuOption) {
	v.SyncWms = false
	poolUpdateCnskuOption.Put(v)
}
