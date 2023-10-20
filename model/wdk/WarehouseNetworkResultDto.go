package wdk

import (
	"sync"
)

// WarehouseNetworkResultDto 结构体
type WarehouseNetworkResultDto struct {
	// 线路信息
	NetworkRouteDtoList []NetworkRouteDto `json:"network_route_dto_list,omitempty" xml:"network_route_dto_list>network_route_dto,omitempty"`
	// 网络类型 WAVE_ARRIVE-波次达。ONE_HOUR-小时达
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 仓商家
	WarehouseMerchantCode string `json:"warehouse_merchant_code,omitempty" xml:"warehouse_merchant_code,omitempty"`
	// 仓编码
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

var poolWarehouseNetworkResultDto = sync.Pool{
	New: func() any {
		return new(WarehouseNetworkResultDto)
	},
}

// GetWarehouseNetworkResultDto() 从对象池中获取WarehouseNetworkResultDto
func GetWarehouseNetworkResultDto() *WarehouseNetworkResultDto {
	return poolWarehouseNetworkResultDto.Get().(*WarehouseNetworkResultDto)
}

// ReleaseWarehouseNetworkResultDto 释放WarehouseNetworkResultDto
func ReleaseWarehouseNetworkResultDto(v *WarehouseNetworkResultDto) {
	v.NetworkRouteDtoList = v.NetworkRouteDtoList[:0]
	v.ServiceType = ""
	v.WarehouseMerchantCode = ""
	v.WarehouseCode = ""
	poolWarehouseNetworkResultDto.Put(v)
}
