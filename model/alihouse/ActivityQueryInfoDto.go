package alihouse

import (
	"sync"
)

// ActivityQueryInfoDto 结构体
type ActivityQueryInfoDto struct {
	// 外部查询ID
	OuterTargetIds []string `json:"outer_target_ids,omitempty" xml:"outer_target_ids>string,omitempty"`
	// 业务域 1：新房 2：二手房 3：租房
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 类型1-公司  2-城市品牌店 3-品牌  4-标准门店  5-签约公司 6-经纪人/置业顾问 7-TP账号
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 查询的场景类型
	SceneType int64 `json:"scene_type,omitempty" xml:"scene_type,omitempty"`
}

var poolActivityQueryInfoDto = sync.Pool{
	New: func() any {
		return new(ActivityQueryInfoDto)
	},
}

// GetActivityQueryInfoDto() 从对象池中获取ActivityQueryInfoDto
func GetActivityQueryInfoDto() *ActivityQueryInfoDto {
	return poolActivityQueryInfoDto.Get().(*ActivityQueryInfoDto)
}

// ReleaseActivityQueryInfoDto 释放ActivityQueryInfoDto
func ReleaseActivityQueryInfoDto(v *ActivityQueryInfoDto) {
	v.OuterTargetIds = v.OuterTargetIds[:0]
	v.BusinessType = 0
	v.Type = 0
	v.SceneType = 0
	poolActivityQueryInfoDto.Put(v)
}
