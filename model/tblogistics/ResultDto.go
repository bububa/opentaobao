package tblogistics

import (
	"sync"
)

// ResultDto 结构体
type ResultDto struct {
	// 返回核销订单列表
	WriteoffOrderList []WriteOffOrderDto `json:"writeoff_order_list,omitempty" xml:"writeoff_order_list>write_off_order_dto,omitempty"`
	// -
	Consign *ConsignDto `json:"consign,omitempty" xml:"consign,omitempty"`
	// 执行结果
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultDto = sync.Pool{
	New: func() any {
		return new(ResultDto)
	},
}

// GetResultDto() 从对象池中获取ResultDto
func GetResultDto() *ResultDto {
	return poolResultDto.Get().(*ResultDto)
}

// ReleaseResultDto 释放ResultDto
func ReleaseResultDto(v *ResultDto) {
	v.WriteoffOrderList = v.WriteoffOrderList[:0]
	v.Consign = nil
	v.Success = false
	poolResultDto.Put(v)
}
