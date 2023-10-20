package nrt

import (
	"sync"
)

// MacallineItemExtDto 结构体
type MacallineItemExtDto struct {
	// 品牌系列ID
	BrandSeriesId string `json:"brand_series_id,omitempty" xml:"brand_series_id,omitempty"`
	// 品牌系列名称
	BrandSeriesName string `json:"brand_series_name,omitempty" xml:"brand_series_name,omitempty"`
	// 计价单位
	ChargeUnit string `json:"charge_unit,omitempty" xml:"charge_unit,omitempty"`
	// 物价员
	Pricer string `json:"pricer,omitempty" xml:"pricer,omitempty"`
	// 辅材
	SecondarySteel string `json:"secondary_steel,omitempty" xml:"secondary_steel,omitempty"`
	// 规格
	Specification string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 规格单位
	SpecificationUnit string `json:"specification_unit,omitempty" xml:"specification_unit,omitempty"`
	// 基材
	Substrate string `json:"substrate,omitempty" xml:"substrate,omitempty"`
	// 饰面
	Veneer string `json:"veneer,omitempty" xml:"veneer,omitempty"`
	// 等级
	Grade int64 `json:"grade,omitempty" xml:"grade,omitempty"`
	// 标签价格类型
	LabelPriceType int64 `json:"label_price_type,omitempty" xml:"label_price_type,omitempty"`
	// 产地
	OriginalLocation *LocationDto `json:"original_location,omitempty" xml:"original_location,omitempty"`
	// 价格类型，1：明码实价，2：明码议价
	PriceType int64 `json:"price_type,omitempty" xml:"price_type,omitempty"`
	// 是否支持退换货
	SupportReturnGoods bool `json:"support_return_goods,omitempty" xml:"support_return_goods,omitempty"`
}

var poolMacallineItemExtDto = sync.Pool{
	New: func() any {
		return new(MacallineItemExtDto)
	},
}

// GetMacallineItemExtDto() 从对象池中获取MacallineItemExtDto
func GetMacallineItemExtDto() *MacallineItemExtDto {
	return poolMacallineItemExtDto.Get().(*MacallineItemExtDto)
}

// ReleaseMacallineItemExtDto 释放MacallineItemExtDto
func ReleaseMacallineItemExtDto(v *MacallineItemExtDto) {
	v.BrandSeriesId = ""
	v.BrandSeriesName = ""
	v.ChargeUnit = ""
	v.Pricer = ""
	v.SecondarySteel = ""
	v.Specification = ""
	v.SpecificationUnit = ""
	v.Substrate = ""
	v.Veneer = ""
	v.Grade = 0
	v.LabelPriceType = 0
	v.OriginalLocation = nil
	v.PriceType = 0
	v.SupportReturnGoods = false
	poolMacallineItemExtDto.Put(v)
}
