package logistic

import (
	"sync"
)

// ExpressModifyAppointTopResponseDto 结构体
type ExpressModifyAppointTopResponseDto struct {
	// 订单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 是否执行成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolExpressModifyAppointTopResponseDto = sync.Pool{
	New: func() any {
		return new(ExpressModifyAppointTopResponseDto)
	},
}

// GetExpressModifyAppointTopResponseDto() 从对象池中获取ExpressModifyAppointTopResponseDto
func GetExpressModifyAppointTopResponseDto() *ExpressModifyAppointTopResponseDto {
	return poolExpressModifyAppointTopResponseDto.Get().(*ExpressModifyAppointTopResponseDto)
}

// ReleaseExpressModifyAppointTopResponseDto 释放ExpressModifyAppointTopResponseDto
func ReleaseExpressModifyAppointTopResponseDto(v *ExpressModifyAppointTopResponseDto) {
	v.OrderCode = ""
	v.Success = false
	poolExpressModifyAppointTopResponseDto.Put(v)
}
