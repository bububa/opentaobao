package alitripmerchant

import (
	"sync"
)

// DerbyVoucherCardCancelDto 结构体
type DerbyVoucherCardCancelDto struct {
	// 订单id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 订单状态code
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 模拟token
	FkToken string `json:"fk_token,omitempty" xml:"fk_token,omitempty"`
}

var poolDerbyVoucherCardCancelDto = sync.Pool{
	New: func() any {
		return new(DerbyVoucherCardCancelDto)
	},
}

// GetDerbyVoucherCardCancelDto() 从对象池中获取DerbyVoucherCardCancelDto
func GetDerbyVoucherCardCancelDto() *DerbyVoucherCardCancelDto {
	return poolDerbyVoucherCardCancelDto.Get().(*DerbyVoucherCardCancelDto)
}

// ReleaseDerbyVoucherCardCancelDto 释放DerbyVoucherCardCancelDto
func ReleaseDerbyVoucherCardCancelDto(v *DerbyVoucherCardCancelDto) {
	v.OrderId = ""
	v.Status = ""
	v.FkToken = ""
	poolDerbyVoucherCardCancelDto.Put(v)
}
