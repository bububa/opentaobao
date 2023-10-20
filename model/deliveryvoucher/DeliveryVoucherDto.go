package deliveryvoucher

import (
	"sync"
)

// DeliveryVoucherDto 结构体
type DeliveryVoucherDto struct {
	// 券ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 券号
	No string `json:"no,omitempty" xml:"no,omitempty"`
}

var poolDeliveryVoucherDto = sync.Pool{
	New: func() any {
		return new(DeliveryVoucherDto)
	},
}

// GetDeliveryVoucherDto() 从对象池中获取DeliveryVoucherDto
func GetDeliveryVoucherDto() *DeliveryVoucherDto {
	return poolDeliveryVoucherDto.Get().(*DeliveryVoucherDto)
}

// ReleaseDeliveryVoucherDto 释放DeliveryVoucherDto
func ReleaseDeliveryVoucherDto(v *DeliveryVoucherDto) {
	v.Id = ""
	v.No = ""
	poolDeliveryVoucherDto.Put(v)
}
