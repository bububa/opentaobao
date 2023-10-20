package tblogistics

import (
	"sync"
)

// ResourceDto 结构体
type ResourceDto struct {
	// 下单凭证
	ResourceRequestId string `json:"resource_request_id,omitempty" xml:"resource_request_id,omitempty"`
	// 资源CODE，用于发货
	ResourceCode string `json:"resource_code,omitempty" xml:"resource_code,omitempty"`
	// 资源名称
	ResourceName string `json:"resource_name,omitempty" xml:"resource_name,omitempty"`
	// 费用明细
	FeeDetail string `json:"fee_detail,omitempty" xml:"fee_detail,omitempty"`
	// 失效原因
	InvalidReason string `json:"invalid_reason,omitempty" xml:"invalid_reason,omitempty"`
	// 扩展
	Features string `json:"features,omitempty" xml:"features,omitempty"`
	// 原价
	OriginaFee int64 `json:"origina_fee,omitempty" xml:"origina_fee,omitempty"`
	// 实付价格
	ActualFee int64 `json:"actual_fee,omitempty" xml:"actual_fee,omitempty"`
	// 是否有效
	Valid bool `json:"valid,omitempty" xml:"valid,omitempty"`
}

var poolResourceDto = sync.Pool{
	New: func() any {
		return new(ResourceDto)
	},
}

// GetResourceDto() 从对象池中获取ResourceDto
func GetResourceDto() *ResourceDto {
	return poolResourceDto.Get().(*ResourceDto)
}

// ReleaseResourceDto 释放ResourceDto
func ReleaseResourceDto(v *ResourceDto) {
	v.ResourceRequestId = ""
	v.ResourceCode = ""
	v.ResourceName = ""
	v.FeeDetail = ""
	v.InvalidReason = ""
	v.Features = ""
	v.OriginaFee = 0
	v.ActualFee = 0
	v.Valid = false
	poolResourceDto.Put(v)
}
