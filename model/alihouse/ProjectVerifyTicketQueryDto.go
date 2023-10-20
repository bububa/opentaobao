package alihouse

import (
	"sync"
)

// ProjectVerifyTicketQueryDto 结构体
type ProjectVerifyTicketQueryDto struct {
	// 核销卷商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolProjectVerifyTicketQueryDto = sync.Pool{
	New: func() any {
		return new(ProjectVerifyTicketQueryDto)
	},
}

// GetProjectVerifyTicketQueryDto() 从对象池中获取ProjectVerifyTicketQueryDto
func GetProjectVerifyTicketQueryDto() *ProjectVerifyTicketQueryDto {
	return poolProjectVerifyTicketQueryDto.Get().(*ProjectVerifyTicketQueryDto)
}

// ReleaseProjectVerifyTicketQueryDto 释放ProjectVerifyTicketQueryDto
func ReleaseProjectVerifyTicketQueryDto(v *ProjectVerifyTicketQueryDto) {
	v.ItemId = 0
	poolProjectVerifyTicketQueryDto.Put(v)
}
