package alihouse

import (
	"sync"
)

// SyncHouseFeaturesDto 结构体
type SyncHouseFeaturesDto struct {
	// 商业活动标
	BusinessActivities []string `json:"business_activities,omitempty" xml:"business_activities>string,omitempty"`
	// 房源列表
	HouseFeaturesList []HouseFeaturesDto `json:"house_features_list,omitempty" xml:"house_features_list>house_features_dto,omitempty"`
	// 操作类型 add-新增，delete-删除
	OperationType string `json:"operation_type,omitempty" xml:"operation_type,omitempty"`
	// 请求批次号
	BatchNumber string `json:"batch_number,omitempty" xml:"batch_number,omitempty"`
	// house-房源类型，store-标准门店
	FeaturesType string `json:"features_type,omitempty" xml:"features_type,omitempty"`
	// 业务类型，1-新房，2-二手房，3-租房，默认为2
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
}

var poolSyncHouseFeaturesDto = sync.Pool{
	New: func() any {
		return new(SyncHouseFeaturesDto)
	},
}

// GetSyncHouseFeaturesDto() 从对象池中获取SyncHouseFeaturesDto
func GetSyncHouseFeaturesDto() *SyncHouseFeaturesDto {
	return poolSyncHouseFeaturesDto.Get().(*SyncHouseFeaturesDto)
}

// ReleaseSyncHouseFeaturesDto 释放SyncHouseFeaturesDto
func ReleaseSyncHouseFeaturesDto(v *SyncHouseFeaturesDto) {
	v.BusinessActivities = v.BusinessActivities[:0]
	v.HouseFeaturesList = v.HouseFeaturesList[:0]
	v.OperationType = ""
	v.BatchNumber = ""
	v.FeaturesType = ""
	v.BusinessType = 0
	poolSyncHouseFeaturesDto.Put(v)
}
