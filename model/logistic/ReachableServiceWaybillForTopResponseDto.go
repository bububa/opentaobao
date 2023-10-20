package logistic

import (
	"sync"
)

// ReachableServiceWaybillForTopResponseDto 结构体
type ReachableServiceWaybillForTopResponseDto struct {
	// 结果列表
	ResultList []ReachableServiceWaybillResponseDto `json:"result_list,omitempty" xml:"result_list>reachable_service_waybill_response_dto,omitempty"`
}

var poolReachableServiceWaybillForTopResponseDto = sync.Pool{
	New: func() any {
		return new(ReachableServiceWaybillForTopResponseDto)
	},
}

// GetReachableServiceWaybillForTopResponseDto() 从对象池中获取ReachableServiceWaybillForTopResponseDto
func GetReachableServiceWaybillForTopResponseDto() *ReachableServiceWaybillForTopResponseDto {
	return poolReachableServiceWaybillForTopResponseDto.Get().(*ReachableServiceWaybillForTopResponseDto)
}

// ReleaseReachableServiceWaybillForTopResponseDto 释放ReachableServiceWaybillForTopResponseDto
func ReleaseReachableServiceWaybillForTopResponseDto(v *ReachableServiceWaybillForTopResponseDto) {
	v.ResultList = v.ResultList[:0]
	poolReachableServiceWaybillForTopResponseDto.Put(v)
}
