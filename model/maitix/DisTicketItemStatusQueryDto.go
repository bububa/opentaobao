package maitix

import (
	"sync"
)

// DisTicketItemStatusQueryDto 结构体
type DisTicketItemStatusQueryDto struct {
	// 票品id列表
	TicketItemIdList []int64 `json:"ticket_item_id_list,omitempty" xml:"ticket_item_id_list>int64,omitempty"`
}

var poolDisTicketItemStatusQueryDto = sync.Pool{
	New: func() any {
		return new(DisTicketItemStatusQueryDto)
	},
}

// GetDisTicketItemStatusQueryDto() 从对象池中获取DisTicketItemStatusQueryDto
func GetDisTicketItemStatusQueryDto() *DisTicketItemStatusQueryDto {
	return poolDisTicketItemStatusQueryDto.Get().(*DisTicketItemStatusQueryDto)
}

// ReleaseDisTicketItemStatusQueryDto 释放DisTicketItemStatusQueryDto
func ReleaseDisTicketItemStatusQueryDto(v *DisTicketItemStatusQueryDto) {
	v.TicketItemIdList = v.TicketItemIdList[:0]
	poolDisTicketItemStatusQueryDto.Put(v)
}
