package maitix

import (
	"sync"
)

// DisPerformStatusQueryDto 结构体
type DisPerformStatusQueryDto struct {
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 是否查询场次对应的票品状态
	QueryTicketItemStatus bool `json:"query_ticket_item_status,omitempty" xml:"query_ticket_item_status,omitempty"`
}

var poolDisPerformStatusQueryDto = sync.Pool{
	New: func() any {
		return new(DisPerformStatusQueryDto)
	},
}

// GetDisPerformStatusQueryDto() 从对象池中获取DisPerformStatusQueryDto
func GetDisPerformStatusQueryDto() *DisPerformStatusQueryDto {
	return poolDisPerformStatusQueryDto.Get().(*DisPerformStatusQueryDto)
}

// ReleaseDisPerformStatusQueryDto 释放DisPerformStatusQueryDto
func ReleaseDisPerformStatusQueryDto(v *DisPerformStatusQueryDto) {
	v.PerformId = 0
	v.QueryTicketItemStatus = false
	poolDisPerformStatusQueryDto.Put(v)
}
