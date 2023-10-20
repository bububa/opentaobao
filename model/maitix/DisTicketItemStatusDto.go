package maitix

import (
	"sync"
)

// DisTicketItemStatusDto 结构体
type DisTicketItemStatusDto struct {
	// 项目id
	ProjectId int64 `json:"project_id,omitempty" xml:"project_id,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 票品id
	TicketItemId int64 `json:"ticket_item_id,omitempty" xml:"ticket_item_id,omitempty"`
	// 票品状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 票品子状态
	SubStatus int64 `json:"sub_status,omitempty" xml:"sub_status,omitempty"`
}

var poolDisTicketItemStatusDto = sync.Pool{
	New: func() any {
		return new(DisTicketItemStatusDto)
	},
}

// GetDisTicketItemStatusDto() 从对象池中获取DisTicketItemStatusDto
func GetDisTicketItemStatusDto() *DisTicketItemStatusDto {
	return poolDisTicketItemStatusDto.Get().(*DisTicketItemStatusDto)
}

// ReleaseDisTicketItemStatusDto 释放DisTicketItemStatusDto
func ReleaseDisTicketItemStatusDto(v *DisTicketItemStatusDto) {
	v.ProjectId = 0
	v.PerformId = 0
	v.TicketItemId = 0
	v.Status = 0
	v.SubStatus = 0
	poolDisTicketItemStatusDto.Put(v)
}
