package icbulogistics

import (
	"sync"
)

// PickupInfoDto 结构体
type PickupInfoDto struct {
	// 备用字段（上门揽收服务商），目前为空
	ServiceProvider string `json:"service_provider,omitempty" xml:"service_provider,omitempty"`
	// 上门揽收类型，warehouse_free_pickup：仓库免费上门揽收，warehouse_paid_pickup：仓库收费上门揽收，provider_paid_pickup：服务商收费上门揽收
	PickupType string `json:"pickup_type,omitempty" xml:"pickup_type,omitempty"`
	// 上门揽收类型名称
	PickupTypeName string `json:"pickup_type_name,omitempty" xml:"pickup_type_name,omitempty"`
	// 能否上门揽收
	CanPickup bool `json:"can_pickup,omitempty" xml:"can_pickup,omitempty"`
}

var poolPickupInfoDto = sync.Pool{
	New: func() any {
		return new(PickupInfoDto)
	},
}

// GetPickupInfoDto() 从对象池中获取PickupInfoDto
func GetPickupInfoDto() *PickupInfoDto {
	return poolPickupInfoDto.Get().(*PickupInfoDto)
}

// ReleasePickupInfoDto 释放PickupInfoDto
func ReleasePickupInfoDto(v *PickupInfoDto) {
	v.ServiceProvider = ""
	v.PickupType = ""
	v.PickupTypeName = ""
	v.CanPickup = false
	poolPickupInfoDto.Put(v)
}
