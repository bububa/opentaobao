package ascp

import (
	"sync"
)

// HiErpCloseDto 结构体
type HiErpCloseDto struct {
	// 履约单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 关单原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 货主id
	OwnerId int64 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
}

var poolHiErpCloseDto = sync.Pool{
	New: func() any {
		return new(HiErpCloseDto)
	},
}

// GetHiErpCloseDto() 从对象池中获取HiErpCloseDto
func GetHiErpCloseDto() *HiErpCloseDto {
	return poolHiErpCloseDto.Get().(*HiErpCloseDto)
}

// ReleaseHiErpCloseDto 释放HiErpCloseDto
func ReleaseHiErpCloseDto(v *HiErpCloseDto) {
	v.OrderCode = ""
	v.Reason = ""
	v.OwnerId = 0
	poolHiErpCloseDto.Put(v)
}
