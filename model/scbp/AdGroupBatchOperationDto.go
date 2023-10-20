package scbp

import (
	"sync"
)

// AdGroupBatchOperationDto 结构体
type AdGroupBatchOperationDto struct {
	// 入参
	AdGroupOperationList []AdGroupOperationDto `json:"ad_group_operation_list,omitempty" xml:"ad_group_operation_list>ad_group_operation_dto,omitempty"`
}

var poolAdGroupBatchOperationDto = sync.Pool{
	New: func() any {
		return new(AdGroupBatchOperationDto)
	},
}

// GetAdGroupBatchOperationDto() 从对象池中获取AdGroupBatchOperationDto
func GetAdGroupBatchOperationDto() *AdGroupBatchOperationDto {
	return poolAdGroupBatchOperationDto.Get().(*AdGroupBatchOperationDto)
}

// ReleaseAdGroupBatchOperationDto 释放AdGroupBatchOperationDto
func ReleaseAdGroupBatchOperationDto(v *AdGroupBatchOperationDto) {
	v.AdGroupOperationList = v.AdGroupOperationList[:0]
	poolAdGroupBatchOperationDto.Put(v)
}
