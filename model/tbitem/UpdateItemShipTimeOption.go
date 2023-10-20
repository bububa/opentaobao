package tbitem

import (
	"sync"
)

// UpdateItemShipTimeOption 结构体
type UpdateItemShipTimeOption struct {
	// 0代表清空匹配的SKU发货时间数据或者商品发货时间数据；1代表：固定发货时间；2代表：相对发货时间
	ShipTimeType int64 `json:"ship_time_type,omitempty" xml:"ship_time_type,omitempty"`
	// 更新类型，默认不填时更新sku，0表示更新sku，1表示更新商品维度，其他值均非法
	UpdateType int64 `json:"update_type,omitempty" xml:"update_type,omitempty"`
}

var poolUpdateItemShipTimeOption = sync.Pool{
	New: func() any {
		return new(UpdateItemShipTimeOption)
	},
}

// GetUpdateItemShipTimeOption() 从对象池中获取UpdateItemShipTimeOption
func GetUpdateItemShipTimeOption() *UpdateItemShipTimeOption {
	return poolUpdateItemShipTimeOption.Get().(*UpdateItemShipTimeOption)
}

// ReleaseUpdateItemShipTimeOption 释放UpdateItemShipTimeOption
func ReleaseUpdateItemShipTimeOption(v *UpdateItemShipTimeOption) {
	v.ShipTimeType = 0
	v.UpdateType = 0
	poolUpdateItemShipTimeOption.Put(v)
}
