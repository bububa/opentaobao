package tttm

import (
	"sync"
)

// SyncPlanDto 结构体
type SyncPlanDto struct {
	// 工单
	SyncWorkDTOs []WorkSyncDto `json:"sync_work_d_t_os,omitempty" xml:"sync_work_d_t_os>work_sync_dto,omitempty"`
	// 计划完成
	FinishTime string `json:"finish_time,omitempty" xml:"finish_time,omitempty"`
	// 订单
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 联系人
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 计划单id
	PlanId string `json:"plan_id,omitempty" xml:"plan_id,omitempty"`
	// 计划开始
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 联系人电话
	ContactTel string `json:"contact_tel,omitempty" xml:"contact_tel,omitempty"`
}

var poolSyncPlanDto = sync.Pool{
	New: func() any {
		return new(SyncPlanDto)
	},
}

// GetSyncPlanDto() 从对象池中获取SyncPlanDto
func GetSyncPlanDto() *SyncPlanDto {
	return poolSyncPlanDto.Get().(*SyncPlanDto)
}

// ReleaseSyncPlanDto 释放SyncPlanDto
func ReleaseSyncPlanDto(v *SyncPlanDto) {
	v.SyncWorkDTOs = v.SyncWorkDTOs[:0]
	v.FinishTime = ""
	v.OrderId = ""
	v.ContactName = ""
	v.PlanId = ""
	v.StartTime = ""
	v.ContactTel = ""
	poolSyncPlanDto.Put(v)
}
