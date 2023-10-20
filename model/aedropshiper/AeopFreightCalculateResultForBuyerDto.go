package aedropshiper

import (
	"sync"
)

// AeopFreightCalculateResultForBuyerDto 结构体
type AeopFreightCalculateResultForBuyerDto struct {
	// 预估运达时效
	EstimatedDeliveryTime string `json:"estimated_delivery_time,omitempty" xml:"estimated_delivery_time,omitempty"`
	// serviceName
	ServiceName string `json:"service_name,omitempty" xml:"service_name,omitempty"`
	// errorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 运费
	Freight *Money `json:"freight,omitempty" xml:"freight,omitempty"`
}

var poolAeopFreightCalculateResultForBuyerDto = sync.Pool{
	New: func() any {
		return new(AeopFreightCalculateResultForBuyerDto)
	},
}

// GetAeopFreightCalculateResultForBuyerDto() 从对象池中获取AeopFreightCalculateResultForBuyerDto
func GetAeopFreightCalculateResultForBuyerDto() *AeopFreightCalculateResultForBuyerDto {
	return poolAeopFreightCalculateResultForBuyerDto.Get().(*AeopFreightCalculateResultForBuyerDto)
}

// ReleaseAeopFreightCalculateResultForBuyerDto 释放AeopFreightCalculateResultForBuyerDto
func ReleaseAeopFreightCalculateResultForBuyerDto(v *AeopFreightCalculateResultForBuyerDto) {
	v.EstimatedDeliveryTime = ""
	v.ServiceName = ""
	v.ErrorCode = 0
	v.Freight = nil
	poolAeopFreightCalculateResultForBuyerDto.Put(v)
}
