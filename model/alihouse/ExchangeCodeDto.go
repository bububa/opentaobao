package alihouse

import (
	"sync"
)

// ExchangeCodeDto 结构体
type ExchangeCodeDto struct {
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 券码
	CouponCode string `json:"coupon_code,omitempty" xml:"coupon_code,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

var poolExchangeCodeDto = sync.Pool{
	New: func() any {
		return new(ExchangeCodeDto)
	},
}

// GetExchangeCodeDto() 从对象池中获取ExchangeCodeDto
func GetExchangeCodeDto() *ExchangeCodeDto {
	return poolExchangeCodeDto.Get().(*ExchangeCodeDto)
}

// ReleaseExchangeCodeDto 释放ExchangeCodeDto
func ReleaseExchangeCodeDto(v *ExchangeCodeDto) {
	v.Remark = ""
	v.CouponCode = ""
	v.StoreId = 0
	poolExchangeCodeDto.Put(v)
}
