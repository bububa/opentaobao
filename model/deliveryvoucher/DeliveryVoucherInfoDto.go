package deliveryvoucher

import (
	"sync"
)

// DeliveryVoucherInfoDto 结构体
type DeliveryVoucherInfoDto struct {
	// 券信息
	Vouchers []DeliveryVoucherDto `json:"vouchers,omitempty" xml:"vouchers>delivery_voucher_dto,omitempty"`
	// 商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolDeliveryVoucherInfoDto = sync.Pool{
	New: func() any {
		return new(DeliveryVoucherInfoDto)
	},
}

// GetDeliveryVoucherInfoDto() 从对象池中获取DeliveryVoucherInfoDto
func GetDeliveryVoucherInfoDto() *DeliveryVoucherInfoDto {
	return poolDeliveryVoucherInfoDto.Get().(*DeliveryVoucherInfoDto)
}

// ReleaseDeliveryVoucherInfoDto 释放DeliveryVoucherInfoDto
func ReleaseDeliveryVoucherInfoDto(v *DeliveryVoucherInfoDto) {
	v.Vouchers = v.Vouchers[:0]
	v.ItemId = 0
	poolDeliveryVoucherInfoDto.Put(v)
}
