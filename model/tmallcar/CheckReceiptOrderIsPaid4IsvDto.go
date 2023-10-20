package tmallcar

import (
	"sync"
)

// CheckReceiptOrderIsPaid4IsvDto 结构体
type CheckReceiptOrderIsPaid4IsvDto struct {
	// 无
	PaidOrderItems []OrderItem4IsvDto `json:"paid_order_items,omitempty" xml:"paid_order_items>order_item4isv_dto,omitempty"`
	// 门店自定义编码
	OuterShopId string `json:"outer_shop_id,omitempty" xml:"outer_shop_id,omitempty"`
	// 工单id
	ReceiptId int64 `json:"receipt_id,omitempty" xml:"receipt_id,omitempty"`
}

var poolCheckReceiptOrderIsPaid4IsvDto = sync.Pool{
	New: func() any {
		return new(CheckReceiptOrderIsPaid4IsvDto)
	},
}

// GetCheckReceiptOrderIsPaid4IsvDto() 从对象池中获取CheckReceiptOrderIsPaid4IsvDto
func GetCheckReceiptOrderIsPaid4IsvDto() *CheckReceiptOrderIsPaid4IsvDto {
	return poolCheckReceiptOrderIsPaid4IsvDto.Get().(*CheckReceiptOrderIsPaid4IsvDto)
}

// ReleaseCheckReceiptOrderIsPaid4IsvDto 释放CheckReceiptOrderIsPaid4IsvDto
func ReleaseCheckReceiptOrderIsPaid4IsvDto(v *CheckReceiptOrderIsPaid4IsvDto) {
	v.PaidOrderItems = v.PaidOrderItems[:0]
	v.OuterShopId = ""
	v.ReceiptId = 0
	poolCheckReceiptOrderIsPaid4IsvDto.Put(v)
}
