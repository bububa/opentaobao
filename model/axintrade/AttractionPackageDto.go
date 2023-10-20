package axintrade

import (
	"sync"
)

// AttractionPackageDto 结构体
type AttractionPackageDto struct {
	// 产品价格
	Prices []ProductPriceDto `json:"prices,omitempty" xml:"prices>product_price_dto,omitempty"`
	// 产品库存
	Inventories []ProductInventoryDto `json:"inventories,omitempty" xml:"inventories>product_inventory_dto,omitempty"`
	// 套餐酒店产品信息
	HotelRates []PackageHotelRateDto `json:"hotel_rates,omitempty" xml:"hotel_rates>package_hotel_rate_dto,omitempty"`
	// 套餐门票产品信息
	TicketInfos []PackageTicketInfoDto `json:"ticket_infos,omitempty" xml:"ticket_infos>package_ticket_info_dto,omitempty"`
	// 餐饮信息
	CateringInfos []PackageCateringInfoDto `json:"catering_infos,omitempty" xml:"catering_infos>package_catering_info_dto,omitempty"`
	// 其他信息
	OtherProducts []PackageOtherProductDto `json:"other_products,omitempty" xml:"other_products>package_other_product_dto,omitempty"`
	// 酒店使用须知
	HotelPolicies []PackageHotelPolicyDto `json:"hotel_policies,omitempty" xml:"hotel_policies>package_hotel_policy_dto,omitempty"`
	// 门票使用须知
	TicketPolicies []PackageTicketPolicyDto `json:"ticket_policies,omitempty" xml:"ticket_policies>package_ticket_policy_dto,omitempty"`
	// 餐饮使用规则
	CateringPolicies []PackageCateringPolicyDto `json:"catering_policies,omitempty" xml:"catering_policies>package_catering_policy_dto,omitempty"`
	// 其他产品使用须知
	OtherPolicies []PackageCateringPolicyDto `json:"other_policies,omitempty" xml:"other_policies>package_catering_policy_dto,omitempty"`
	// 套餐名称
	PackageName string `json:"package_name,omitempty" xml:"package_name,omitempty"`
	// 套餐描述
	PackageDesc string `json:"package_desc,omitempty" xml:"package_desc,omitempty"`
	// 费用不包含
	CostNotInclude string `json:"cost_not_include,omitempty" xml:"cost_not_include,omitempty"`
	// 酒景套餐ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 退改规则
	RefundPolicy *PackageRefundPolicyDto `json:"refund_policy,omitempty" xml:"refund_policy,omitempty"`
	// 预定规则
	BookingPolicy *BookingPolicyDto `json:"booking_policy,omitempty" xml:"booking_policy,omitempty"`
	// 限购规则
	LimitPolicy *LimitPolicyDto `json:"limit_policy,omitempty" xml:"limit_policy,omitempty"`
	// 游客规则
	TouristPolicy *TouristPolicyDto `json:"tourist_policy,omitempty" xml:"tourist_policy,omitempty"`
	// 套餐位置信息
	AreaInfo *ProductAreaInfoDto `json:"area_info,omitempty" xml:"area_info,omitempty"`
}

var poolAttractionPackageDto = sync.Pool{
	New: func() any {
		return new(AttractionPackageDto)
	},
}

// GetAttractionPackageDto() 从对象池中获取AttractionPackageDto
func GetAttractionPackageDto() *AttractionPackageDto {
	return poolAttractionPackageDto.Get().(*AttractionPackageDto)
}

// ReleaseAttractionPackageDto 释放AttractionPackageDto
func ReleaseAttractionPackageDto(v *AttractionPackageDto) {
	v.Prices = v.Prices[:0]
	v.Inventories = v.Inventories[:0]
	v.HotelRates = v.HotelRates[:0]
	v.TicketInfos = v.TicketInfos[:0]
	v.CateringInfos = v.CateringInfos[:0]
	v.OtherProducts = v.OtherProducts[:0]
	v.HotelPolicies = v.HotelPolicies[:0]
	v.TicketPolicies = v.TicketPolicies[:0]
	v.CateringPolicies = v.CateringPolicies[:0]
	v.OtherPolicies = v.OtherPolicies[:0]
	v.PackageName = ""
	v.PackageDesc = ""
	v.CostNotInclude = ""
	v.ProductId = 0
	v.RefundPolicy = nil
	v.BookingPolicy = nil
	v.LimitPolicy = nil
	v.TouristPolicy = nil
	v.AreaInfo = nil
	poolAttractionPackageDto.Put(v)
}
