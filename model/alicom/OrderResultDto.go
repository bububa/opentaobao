package alicom

import (
	"sync"
)

// OrderResultDto 结构体
type OrderResultDto struct {
	// 结果描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 天猫交易订单号
	OrderNo int64 `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 商家处理结果是否成功标志
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolOrderResultDto = sync.Pool{
	New: func() any {
		return new(OrderResultDto)
	},
}

// GetOrderResultDto() 从对象池中获取OrderResultDto
func GetOrderResultDto() *OrderResultDto {
	return poolOrderResultDto.Get().(*OrderResultDto)
}

// ReleaseOrderResultDto 释放OrderResultDto
func ReleaseOrderResultDto(v *OrderResultDto) {
	v.Desc = ""
	v.ResultCode = ""
	v.OrderNo = 0
	v.Success = false
	poolOrderResultDto.Put(v)
}
