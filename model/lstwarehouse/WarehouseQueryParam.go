package lstwarehouse

import (
	"sync"
)

// WarehouseQueryParam 结构体
type WarehouseQueryParam struct {
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页最大记录数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
}

var poolWarehouseQueryParam = sync.Pool{
	New: func() any {
		return new(WarehouseQueryParam)
	},
}

// GetWarehouseQueryParam() 从对象池中获取WarehouseQueryParam
func GetWarehouseQueryParam() *WarehouseQueryParam {
	return poolWarehouseQueryParam.Get().(*WarehouseQueryParam)
}

// ReleaseWarehouseQueryParam 释放WarehouseQueryParam
func ReleaseWarehouseQueryParam(v *WarehouseQueryParam) {
	v.Page = 0
	v.Size = 0
	poolWarehouseQueryParam.Put(v)
}
