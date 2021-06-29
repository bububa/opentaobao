package alihouse

// ProjectPreSalePermitDTO 
type ProjectPreSalePermitDTO struct {
    // 公示时间
    PublicityTime   string `json:"publicity_time,omitempty" xml:"publicity_time,omitempty"`
    // 总价区间
    TotalPriceRange   string `json:"total_price_range,omitempty" xml:"total_price_range,omitempty"`
    // 单价区间
    UnitPriceRange   string `json:"unit_price_range,omitempty" xml:"unit_price_range,omitempty"`
    // 均价分
    AveragePrice   string `json:"average_price,omitempty" xml:"average_price,omitempty"`
    // 登记结束时间
    RegistrationEndTime   string `json:"registration_end_time,omitempty" xml:"registration_end_time,omitempty"`
    // 登记开始时间
    RegistrationStartTime   string `json:"registration_start_time,omitempty" xml:"registration_start_time,omitempty"`
    // 住宅拟售价格(分)
    ProposedSalePrice   string `json:"proposed_sale_price,omitempty" xml:"proposed_sale_price,omitempty"`
    // 销售面积
    SalesArea   int64 `json:"sales_area,omitempty" xml:"sales_area,omitempty"`
    // 销售套数
    SalesSets   int64 `json:"sales_sets,omitempty" xml:"sales_sets,omitempty"`
    // 使用年限
    ServiceLife   string `json:"service_life,omitempty" xml:"service_life,omitempty"`
    // 预售部位
    PreSalePosition   string `json:"pre_sale_position,omitempty" xml:"pre_sale_position,omitempty"`
    // 房屋/土地用途
    LandHousingUse   string `json:"land_housing_use,omitempty" xml:"land_housing_use,omitempty"`
    // 准许销售面积
    PermittedSalesArea   string `json:"permitted_sales_area,omitempty" xml:"permitted_sales_area,omitempty"`
    // 发证日期
    CertificationDate   string `json:"certification_date,omitempty" xml:"certification_date,omitempty"`
    // 预售证编号
    PreSaleLicenseNumber   string `json:"pre_sale_license_number,omitempty" xml:"pre_sale_license_number,omitempty"`
    // 外部预售证id
    OuterPermitId   string `json:"outer_permit_id,omitempty" xml:"outer_permit_id,omitempty"`
    // 外部楼盘ID
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // 预售楼幢
    PreSaleBuild   string `json:"pre_sale_build,omitempty" xml:"pre_sale_build,omitempty"`
}
