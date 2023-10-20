package ascpffo

import (
	"sync"
)

// ReturnOrderQueryDto 结构体
type ReturnOrderQueryDto struct {
	// 退供单号
	ReturnOrderNo string `json:"return_order_no,omitempty" xml:"return_order_no,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

var poolReturnOrderQueryDto = sync.Pool{
	New: func() any {
		return new(ReturnOrderQueryDto)
	},
}

// GetReturnOrderQueryDto() 从对象池中获取ReturnOrderQueryDto
func GetReturnOrderQueryDto() *ReturnOrderQueryDto {
	return poolReturnOrderQueryDto.Get().(*ReturnOrderQueryDto)
}

// ReleaseReturnOrderQueryDto 释放ReturnOrderQueryDto
func ReleaseReturnOrderQueryDto(v *ReturnOrderQueryDto) {
	v.ReturnOrderNo = ""
	v.BizType = 0
	poolReturnOrderQueryDto.Put(v)
}
