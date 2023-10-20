package fenxiao

import (
	"sync"
)

// AddCnskuOption 结构体
type AddCnskuOption struct {
	// 是否同步到wms, 为空时默认下发
	SyncWms bool `json:"sync_wms,omitempty" xml:"sync_wms,omitempty"`
}

var poolAddCnskuOption = sync.Pool{
	New: func() any {
		return new(AddCnskuOption)
	},
}

// GetAddCnskuOption() 从对象池中获取AddCnskuOption
func GetAddCnskuOption() *AddCnskuOption {
	return poolAddCnskuOption.Get().(*AddCnskuOption)
}

// ReleaseAddCnskuOption 释放AddCnskuOption
func ReleaseAddCnskuOption(v *AddCnskuOption) {
	v.SyncWms = false
	poolAddCnskuOption.Put(v)
}
