package tttm

import (
	"sync"
)

// WorkSyncDto 结构体
type WorkSyncDto struct {
	// 生产节点
	ProduceLinks []string `json:"produce_links,omitempty" xml:"produce_links>string,omitempty"`
	// 生产进度
	SyncProduceDTOs []ProduceSyncDto `json:"sync_produce_d_t_os,omitempty" xml:"sync_produce_d_t_os>produce_sync_dto,omitempty"`
	// 工单ID
	WorkId string `json:"work_id,omitempty" xml:"work_id,omitempty"`
	// 货品
	ProductCode string `json:"product_code,omitempty" xml:"product_code,omitempty"`
	// 计划产量
	PlanNum string `json:"plan_num,omitempty" xml:"plan_num,omitempty"`
	// 工单开始
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 工单结束
	FinishTime string `json:"finish_time,omitempty" xml:"finish_time,omitempty"`
	// 生产状态
	ProduceStatus int64 `json:"produce_status,omitempty" xml:"produce_status,omitempty"`
}

var poolWorkSyncDto = sync.Pool{
	New: func() any {
		return new(WorkSyncDto)
	},
}

// GetWorkSyncDto() 从对象池中获取WorkSyncDto
func GetWorkSyncDto() *WorkSyncDto {
	return poolWorkSyncDto.Get().(*WorkSyncDto)
}

// ReleaseWorkSyncDto 释放WorkSyncDto
func ReleaseWorkSyncDto(v *WorkSyncDto) {
	v.ProduceLinks = v.ProduceLinks[:0]
	v.SyncProduceDTOs = v.SyncProduceDTOs[:0]
	v.WorkId = ""
	v.ProductCode = ""
	v.PlanNum = ""
	v.StartTime = ""
	v.FinishTime = ""
	v.ProduceStatus = 0
	poolWorkSyncDto.Put(v)
}
