package logistic

import (
	"sync"
)

// RoutingReachableBatchRequestDto 结构体
type RoutingReachableBatchRequestDto struct {
	// 收发地址和服务列表
	AddressAndServiceList []ReachableAddressAndServiceDto `json:"address_and_service_list,omitempty" xml:"address_and_service_list>reachable_address_and_service_dto,omitempty"`
	// 快递公司code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 发件揽收网点
	SendBranchCode string `json:"send_branch_code,omitempty" xml:"send_branch_code,omitempty"`
}

var poolRoutingReachableBatchRequestDto = sync.Pool{
	New: func() any {
		return new(RoutingReachableBatchRequestDto)
	},
}

// GetRoutingReachableBatchRequestDto() 从对象池中获取RoutingReachableBatchRequestDto
func GetRoutingReachableBatchRequestDto() *RoutingReachableBatchRequestDto {
	return poolRoutingReachableBatchRequestDto.Get().(*RoutingReachableBatchRequestDto)
}

// ReleaseRoutingReachableBatchRequestDto 释放RoutingReachableBatchRequestDto
func ReleaseRoutingReachableBatchRequestDto(v *RoutingReachableBatchRequestDto) {
	v.AddressAndServiceList = v.AddressAndServiceList[:0]
	v.CpCode = ""
	v.SendBranchCode = ""
	poolRoutingReachableBatchRequestDto.Put(v)
}
