package alihealthpw

import (
	"sync"
)

// PendingListDto 结构体
type PendingListDto struct {
	// 申请状态
	ApplyStatus string `json:"apply_status,omitempty" xml:"apply_status,omitempty"`
	// 申请类型
	ApplyType string `json:"apply_type,omitempty" xml:"apply_type,omitempty"`
	// 申请时间
	ApplyDate string `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	// 唯一编码
	ApplyUniqueCode string `json:"apply_unique_code,omitempty" xml:"apply_unique_code,omitempty"`
}

var poolPendingListDto = sync.Pool{
	New: func() any {
		return new(PendingListDto)
	},
}

// GetPendingListDto() 从对象池中获取PendingListDto
func GetPendingListDto() *PendingListDto {
	return poolPendingListDto.Get().(*PendingListDto)
}

// ReleasePendingListDto 释放PendingListDto
func ReleasePendingListDto(v *PendingListDto) {
	v.ApplyStatus = ""
	v.ApplyType = ""
	v.ApplyDate = ""
	v.ApplyUniqueCode = ""
	poolPendingListDto.Put(v)
}
