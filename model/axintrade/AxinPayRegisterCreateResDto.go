package axintrade

import (
	"sync"
)

// AxinPayRegisterCreateResDto 结构体
type AxinPayRegisterCreateResDto struct {
	// 支付宝返回的申请单号
	ApplyOrderId string `json:"apply_order_id,omitempty" xml:"apply_order_id,omitempty"`
	// 支付平台申请单号
	PayRegisterOrderNo string `json:"pay_register_order_no,omitempty" xml:"pay_register_order_no,omitempty"`
}

var poolAxinPayRegisterCreateResDto = sync.Pool{
	New: func() any {
		return new(AxinPayRegisterCreateResDto)
	},
}

// GetAxinPayRegisterCreateResDto() 从对象池中获取AxinPayRegisterCreateResDto
func GetAxinPayRegisterCreateResDto() *AxinPayRegisterCreateResDto {
	return poolAxinPayRegisterCreateResDto.Get().(*AxinPayRegisterCreateResDto)
}

// ReleaseAxinPayRegisterCreateResDto 释放AxinPayRegisterCreateResDto
func ReleaseAxinPayRegisterCreateResDto(v *AxinPayRegisterCreateResDto) {
	v.ApplyOrderId = ""
	v.PayRegisterOrderNo = ""
	poolAxinPayRegisterCreateResDto.Put(v)
}
