package tmallchannel

import (
	"sync"
)

// TopPurchasePayOrderDto 结构体
type TopPurchasePayOrderDto struct {
	// 支付流水编号
	PayOrderId string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	// 支付金额
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
}

var poolTopPurchasePayOrderDto = sync.Pool{
	New: func() any {
		return new(TopPurchasePayOrderDto)
	},
}

// GetTopPurchasePayOrderDto() 从对象池中获取TopPurchasePayOrderDto
func GetTopPurchasePayOrderDto() *TopPurchasePayOrderDto {
	return poolTopPurchasePayOrderDto.Get().(*TopPurchasePayOrderDto)
}

// ReleaseTopPurchasePayOrderDto 释放TopPurchasePayOrderDto
func ReleaseTopPurchasePayOrderDto(v *TopPurchasePayOrderDto) {
	v.PayOrderId = ""
	v.PayFee = 0
	poolTopPurchasePayOrderDto.Put(v)
}
