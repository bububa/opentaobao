package legalsuit

import (
	"sync"
)

// StandpointRefDto 结构体
type StandpointRefDto struct {
	// 口径标签
	StandpointLabels []string `json:"standpoint_labels,omitempty" xml:"standpoint_labels>string,omitempty"`
	// options
	Options []Option `json:"options,omitempty" xml:"options>option,omitempty"`
	// table信息
	TableStr []string `json:"table_str,omitempty" xml:"table_str>string,omitempty"`
	// 观点描述
	StandpointDesc string `json:"standpoint_desc,omitempty" xml:"standpoint_desc,omitempty"`
	// 口径
	DefenseCaliber string `json:"defense_caliber,omitempty" xml:"defense_caliber,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 场景结构
	ScenesStrucct string `json:"scenes_strucct,omitempty" xml:"scenes_strucct,omitempty"`
	// 场景名称
	ScenesName string `json:"scenes_name,omitempty" xml:"scenes_name,omitempty"`
	// 创建者
	CreaterWorkerNo string `json:"creater_worker_no,omitempty" xml:"creater_worker_no,omitempty"`
	// label
	Label string `json:"label,omitempty" xml:"label,omitempty"`
	// tableSchema
	TableSchema string `json:"table_schema,omitempty" xml:"table_schema,omitempty"`
	// 关联观点来源
	RefrenceType string `json:"refrence_type,omitempty" xml:"refrence_type,omitempty"`
	// 场景id
	ScenesId int64 `json:"scenes_id,omitempty" xml:"scenes_id,omitempty"`
	// 关联id
	RefrenceId int64 `json:"refrence_id,omitempty" xml:"refrence_id,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolStandpointRefDto = sync.Pool{
	New: func() any {
		return new(StandpointRefDto)
	},
}

// GetStandpointRefDto() 从对象池中获取StandpointRefDto
func GetStandpointRefDto() *StandpointRefDto {
	return poolStandpointRefDto.Get().(*StandpointRefDto)
}

// ReleaseStandpointRefDto 释放StandpointRefDto
func ReleaseStandpointRefDto(v *StandpointRefDto) {
	v.StandpointLabels = v.StandpointLabels[:0]
	v.Options = v.Options[:0]
	v.TableStr = v.TableStr[:0]
	v.StandpointDesc = ""
	v.DefenseCaliber = ""
	v.GmtCreate = ""
	v.ScenesStrucct = ""
	v.ScenesName = ""
	v.CreaterWorkerNo = ""
	v.Label = ""
	v.TableSchema = ""
	v.RefrenceType = ""
	v.ScenesId = 0
	v.RefrenceId = 0
	v.Id = 0
	poolStandpointRefDto.Put(v)
}
