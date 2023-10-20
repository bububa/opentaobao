package tbk

import (
	"sync"
)

// PlatformSpecialShareInfoDto 结构体
type PlatformSpecialShareInfoDto struct {
	// 内容专项服务费比率
	ContentTechServiceRate float64 `json:"content_tech_service_rate,omitempty" xml:"content_tech_service_rate,omitempty"`
	// 预估内容专项服务费
	ContentTechServicePreFee float64 `json:"content_tech_service_pre_fee,omitempty" xml:"content_tech_service_pre_fee,omitempty"`
	// 结算内容专项服务费
	ContentTechServiceFee float64 `json:"content_tech_service_fee,omitempty" xml:"content_tech_service_fee,omitempty"`
	// 流量专项服务费比率（默认无，限定开放）
	TrafficTechServiceRate float64 `json:"traffic_tech_service_rate,omitempty" xml:"traffic_tech_service_rate,omitempty"`
	// 预估流量专项服务费（默认无，限定开放）
	TrafficTechServicePreFee float64 `json:"traffic_tech_service_pre_fee,omitempty" xml:"traffic_tech_service_pre_fee,omitempty"`
	// 结算流量专项服务费（默认无，限定开放）
	TrafficTechServiceFee float64 `json:"traffic_tech_service_fee,omitempty" xml:"traffic_tech_service_fee,omitempty"`
}

var poolPlatformSpecialShareInfoDto = sync.Pool{
	New: func() any {
		return new(PlatformSpecialShareInfoDto)
	},
}

// GetPlatformSpecialShareInfoDto() 从对象池中获取PlatformSpecialShareInfoDto
func GetPlatformSpecialShareInfoDto() *PlatformSpecialShareInfoDto {
	return poolPlatformSpecialShareInfoDto.Get().(*PlatformSpecialShareInfoDto)
}

// ReleasePlatformSpecialShareInfoDto 释放PlatformSpecialShareInfoDto
func ReleasePlatformSpecialShareInfoDto(v *PlatformSpecialShareInfoDto) {
	v.ContentTechServiceRate = 0
	v.ContentTechServicePreFee = 0
	v.ContentTechServiceFee = 0
	v.TrafficTechServiceRate = 0
	v.TrafficTechServicePreFee = 0
	v.TrafficTechServiceFee = 0
	poolPlatformSpecialShareInfoDto.Put(v)
}
