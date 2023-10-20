package gameact

import (
	"sync"
)

// UpdateDeliveryAddressVo 结构体
type UpdateDeliveryAddressVo struct {
	// 是否成功更新或确认地址
	UpdateAddress bool `json:"update_address,omitempty" xml:"update_address,omitempty"`
}

var poolUpdateDeliveryAddressVo = sync.Pool{
	New: func() any {
		return new(UpdateDeliveryAddressVo)
	},
}

// GetUpdateDeliveryAddressVo() 从对象池中获取UpdateDeliveryAddressVo
func GetUpdateDeliveryAddressVo() *UpdateDeliveryAddressVo {
	return poolUpdateDeliveryAddressVo.Get().(*UpdateDeliveryAddressVo)
}

// ReleaseUpdateDeliveryAddressVo 释放UpdateDeliveryAddressVo
func ReleaseUpdateDeliveryAddressVo(v *UpdateDeliveryAddressVo) {
	v.UpdateAddress = false
	poolUpdateDeliveryAddressVo.Put(v)
}
