package alihouse

import (
	"sync"
)

// PreSalePermitDto 结构体
type PreSalePermitDto struct {
	// 预售楼幢外部id列表
	OuterBuildingIds []string `json:"outer_building_ids,omitempty" xml:"outer_building_ids>string,omitempty"`
	// 公示时间
	PublicityTime string `json:"publicity_time,omitempty" xml:"publicity_time,omitempty"`
	// 总价区间
	TotalPriceRange string `json:"total_price_range,omitempty" xml:"total_price_range,omitempty"`
	// 单价区间
	UnitPriceRange string `json:"unit_price_range,omitempty" xml:"unit_price_range,omitempty"`
	// 均价分
	AveragePrice string `json:"average_price,omitempty" xml:"average_price,omitempty"`
	// 登记结束时间
	RegistrationEndTime string `json:"registration_end_time,omitempty" xml:"registration_end_time,omitempty"`
	// 登记开始时间
	RegistrationStartTime string `json:"registration_start_time,omitempty" xml:"registration_start_time,omitempty"`
	// 住宅拟售价格(分)
	ProposedSalePrice string `json:"proposed_sale_price,omitempty" xml:"proposed_sale_price,omitempty"`
	// 使用年限
	ServiceLife string `json:"service_life,omitempty" xml:"service_life,omitempty"`
	// 预售部位
	PreSalePosition string `json:"pre_sale_position,omitempty" xml:"pre_sale_position,omitempty"`
	// 房屋/土地用途
	LandHousingUse string `json:"land_housing_use,omitempty" xml:"land_housing_use,omitempty"`
	// 准许销售面积
	PermittedSalesArea string `json:"permitted_sales_area,omitempty" xml:"permitted_sales_area,omitempty"`
	// 发证日期
	CertificationDate string `json:"certification_date,omitempty" xml:"certification_date,omitempty"`
	// 预售证编号
	PreSaleLicenseNumber string `json:"pre_sale_license_number,omitempty" xml:"pre_sale_license_number,omitempty"`
	// 外部预售证id
	OuterPermitId string `json:"outer_permit_id,omitempty" xml:"outer_permit_id,omitempty"`
	// 外部楼盘ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 外部项目店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 销售面积
	SalesArea int64 `json:"sales_area,omitempty" xml:"sales_area,omitempty"`
	// 销售套数
	SalesSets int64 `json:"sales_sets,omitempty" xml:"sales_sets,omitempty"`
}

var poolPreSalePermitDto = sync.Pool{
	New: func() any {
		return new(PreSalePermitDto)
	},
}

// GetPreSalePermitDto() 从对象池中获取PreSalePermitDto
func GetPreSalePermitDto() *PreSalePermitDto {
	return poolPreSalePermitDto.Get().(*PreSalePermitDto)
}

// ReleasePreSalePermitDto 释放PreSalePermitDto
func ReleasePreSalePermitDto(v *PreSalePermitDto) {
	v.OuterBuildingIds = v.OuterBuildingIds[:0]
	v.PublicityTime = ""
	v.TotalPriceRange = ""
	v.UnitPriceRange = ""
	v.AveragePrice = ""
	v.RegistrationEndTime = ""
	v.RegistrationStartTime = ""
	v.ProposedSalePrice = ""
	v.ServiceLife = ""
	v.PreSalePosition = ""
	v.LandHousingUse = ""
	v.PermittedSalesArea = ""
	v.CertificationDate = ""
	v.PreSaleLicenseNumber = ""
	v.OuterPermitId = ""
	v.OuterId = ""
	v.OuterStoreId = ""
	v.SalesArea = 0
	v.SalesSets = 0
	poolPreSalePermitDto.Put(v)
}
