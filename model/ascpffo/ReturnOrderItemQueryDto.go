package ascpffo

import (
	"sync"
)

// ReturnOrderItemQueryDto 结构体
type ReturnOrderItemQueryDto struct {
	// 退供单号
	ReturnOrderNo string `json:"return_order_no,omitempty" xml:"return_order_no,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolReturnOrderItemQueryDto = sync.Pool{
	New: func() any {
		return new(ReturnOrderItemQueryDto)
	},
}

// GetReturnOrderItemQueryDto() 从对象池中获取ReturnOrderItemQueryDto
func GetReturnOrderItemQueryDto() *ReturnOrderItemQueryDto {
	return poolReturnOrderItemQueryDto.Get().(*ReturnOrderItemQueryDto)
}

// ReleaseReturnOrderItemQueryDto 释放ReturnOrderItemQueryDto
func ReleaseReturnOrderItemQueryDto(v *ReturnOrderItemQueryDto) {
	v.ReturnOrderNo = ""
	v.BizType = 0
	v.PageIndex = 0
	v.PageSize = 0
	poolReturnOrderItemQueryDto.Put(v)
}
