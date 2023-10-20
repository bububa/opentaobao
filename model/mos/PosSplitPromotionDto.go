package mos

import (
	"sync"
)

// PosSplitPromotionDto 结构体
type PosSplitPromotionDto struct {
	// 券码
	CouponCode string `json:"coupon_code,omitempty" xml:"coupon_code,omitempty"`
	// 结算码
	SettleCode string `json:"settle_code,omitempty" xml:"settle_code,omitempty"`
	// 该支付方式支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 大支付方式
	Payment string `json:"payment,omitempty" xml:"payment,omitempty"`
	// 小支付方式
	SubPayment string `json:"sub_payment,omitempty" xml:"sub_payment,omitempty"`
	// 是否压卡
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
	// 商品行号，必须
	GoodsLineNo int64 `json:"goods_line_no,omitempty" xml:"goods_line_no,omitempty"`
	// 支付行号,必须
	PaymentLineNo int64 `json:"payment_line_no,omitempty" xml:"payment_line_no,omitempty"`
	// 分摊金额，必须
	SplitAmount int64 `json:"split_amount,omitempty" xml:"split_amount,omitempty"`
	// 1:支付，2:优惠
	PayType int64 `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
}

var poolPosSplitPromotionDto = sync.Pool{
	New: func() any {
		return new(PosSplitPromotionDto)
	},
}

// GetPosSplitPromotionDto() 从对象池中获取PosSplitPromotionDto
func GetPosSplitPromotionDto() *PosSplitPromotionDto {
	return poolPosSplitPromotionDto.Get().(*PosSplitPromotionDto)
}

// ReleasePosSplitPromotionDto 释放PosSplitPromotionDto
func ReleasePosSplitPromotionDto(v *PosSplitPromotionDto) {
	v.CouponCode = ""
	v.SettleCode = ""
	v.PayTime = ""
	v.Payment = ""
	v.SubPayment = ""
	v.ExtendParams = ""
	v.GoodsLineNo = 0
	v.PaymentLineNo = 0
	v.SplitAmount = 0
	v.PayType = 0
	poolPosSplitPromotionDto.Put(v)
}
