package idleisv

import (
	"sync"
)

// AppraiseIsvOrderCloseDto 结构体
type AppraiseIsvOrderCloseDto struct {
	// 关闭订单原因
	CloseReason string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// 订单id
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolAppraiseIsvOrderCloseDto = sync.Pool{
	New: func() any {
		return new(AppraiseIsvOrderCloseDto)
	},
}

// GetAppraiseIsvOrderCloseDto() 从对象池中获取AppraiseIsvOrderCloseDto
func GetAppraiseIsvOrderCloseDto() *AppraiseIsvOrderCloseDto {
	return poolAppraiseIsvOrderCloseDto.Get().(*AppraiseIsvOrderCloseDto)
}

// ReleaseAppraiseIsvOrderCloseDto 释放AppraiseIsvOrderCloseDto
func ReleaseAppraiseIsvOrderCloseDto(v *AppraiseIsvOrderCloseDto) {
	v.CloseReason = ""
	v.BizOrderId = 0
	poolAppraiseIsvOrderCloseDto.Put(v)
}
