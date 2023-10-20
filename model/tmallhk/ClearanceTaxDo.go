package tmallhk

import (
	"sync"
)

// ClearanceTaxDo 结构体
type ClearanceTaxDo struct {
	// 海关税收编码
	Hscode string `json:"hscode,omitempty" xml:"hscode,omitempty"`
	// 第一数量，最多4位小数
	FirstQuantity string `json:"first_quantity,omitempty" xml:"first_quantity,omitempty"`
	// 第二数量，最多4位小数
	SecondQuantity string `json:"second_quantity,omitempty" xml:"second_quantity,omitempty"`
	// 第一单位，单位编码
	FirstUnit string `json:"first_unit,omitempty" xml:"first_unit,omitempty"`
	// 第二单位，单位编码
	SecondUnit string `json:"second_unit,omitempty" xml:"second_unit,omitempty"`
	// 消费税，主&amp;子
	ExciseDutyFee int64 `json:"excise_duty_fee,omitempty" xml:"excise_duty_fee,omitempty"`
	// 税费，子订单
	OrderLineTotalTaxFee int64 `json:"order_line_total_tax_fee,omitempty" xml:"order_line_total_tax_fee,omitempty"`
	// 邮费，主&amp;子
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 关税，主&amp;子
	CustomDutyFee int64 `json:"custom_duty_fee,omitempty" xml:"custom_duty_fee,omitempty"`
	// 给海关的税费中的增值税，主&amp;子
	TariffVatFee int64 `json:"tariff_vat_fee,omitempty" xml:"tariff_vat_fee,omitempty"`
	// 给海关的税费中的消费税，主&amp;子
	TariffExciseFee int64 `json:"tariff_excise_fee,omitempty" xml:"tariff_excise_fee,omitempty"`
	// 增值税，主&amp;子
	VatFee int64 `json:"vat_fee,omitempty" xml:"vat_fee,omitempty"`
	// 优惠，主&amp;子
	CustomsCouponFee int64 `json:"customs_coupon_fee,omitempty" xml:"customs_coupon_fee,omitempty"`
	// 包税包不仅，给海关的税费，主&amp;子
	TariffFee int64 `json:"tariff_fee,omitempty" xml:"tariff_fee,omitempty"`
	// 完税价，子订单
	CustomsSubTotalFee int64 `json:"customs_sub_total_fee,omitempty" xml:"customs_sub_total_fee,omitempty"`
	// 包税包不仅，给海关的税费中的关税，主&amp;子
	TariffCustomFee int64 `json:"tariff_custom_fee,omitempty" xml:"tariff_custom_fee,omitempty"`
	// 保费，主&amp;子
	CustomsInsuranceFee int64 `json:"customs_insurance_fee,omitempty" xml:"customs_insurance_fee,omitempty"`
	// 总税费，主订单
	OrderTotalTaxFee int64 `json:"order_total_tax_fee,omitempty" xml:"order_total_tax_fee,omitempty"`
	// 完税价，主订单
	CustomsTotalFee int64 `json:"customs_total_fee,omitempty" xml:"customs_total_fee,omitempty"`
}

var poolClearanceTaxDo = sync.Pool{
	New: func() any {
		return new(ClearanceTaxDo)
	},
}

// GetClearanceTaxDo() 从对象池中获取ClearanceTaxDo
func GetClearanceTaxDo() *ClearanceTaxDo {
	return poolClearanceTaxDo.Get().(*ClearanceTaxDo)
}

// ReleaseClearanceTaxDo 释放ClearanceTaxDo
func ReleaseClearanceTaxDo(v *ClearanceTaxDo) {
	v.Hscode = ""
	v.FirstQuantity = ""
	v.SecondQuantity = ""
	v.FirstUnit = ""
	v.SecondUnit = ""
	v.ExciseDutyFee = 0
	v.OrderLineTotalTaxFee = 0
	v.PostFee = 0
	v.CustomDutyFee = 0
	v.TariffVatFee = 0
	v.TariffExciseFee = 0
	v.VatFee = 0
	v.CustomsCouponFee = 0
	v.TariffFee = 0
	v.CustomsSubTotalFee = 0
	v.TariffCustomFee = 0
	v.CustomsInsuranceFee = 0
	v.OrderTotalTaxFee = 0
	v.CustomsTotalFee = 0
	poolClearanceTaxDo.Put(v)
}
