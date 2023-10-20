package alihouse

import (
	"sync"
)

// SyncBrokerPointsDto 结构体
type SyncBrokerPointsDto struct {
	// 经纪人积分因素列表
	PointsFactors []PointsFactorDto `json:"points_factors,omitempty" xml:"points_factors>points_factor_dto,omitempty"`
	// 外部经纪人ID
	OuterConsultantId string `json:"outer_consultant_id,omitempty" xml:"outer_consultant_id,omitempty"`
	// 业务数据归属时间
	BusinessTime string `json:"business_time,omitempty" xml:"business_time,omitempty"`
	// 0默认不是测试，1是测试数据
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolSyncBrokerPointsDto = sync.Pool{
	New: func() any {
		return new(SyncBrokerPointsDto)
	},
}

// GetSyncBrokerPointsDto() 从对象池中获取SyncBrokerPointsDto
func GetSyncBrokerPointsDto() *SyncBrokerPointsDto {
	return poolSyncBrokerPointsDto.Get().(*SyncBrokerPointsDto)
}

// ReleaseSyncBrokerPointsDto 释放SyncBrokerPointsDto
func ReleaseSyncBrokerPointsDto(v *SyncBrokerPointsDto) {
	v.PointsFactors = v.PointsFactors[:0]
	v.OuterConsultantId = ""
	v.BusinessTime = ""
	v.IsTest = 0
	poolSyncBrokerPointsDto.Put(v)
}
