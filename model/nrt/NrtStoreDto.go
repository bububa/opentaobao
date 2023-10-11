package nrt

// NrtStoreDto 结构体
type NrtStoreDto struct {
	// 摊位编号
	StallCode string `json:"stall_code,omitempty" xml:"stall_code,omitempty"`
	// 业务标识
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 详细地址
	ShortAddress string `json:"short_address,omitempty" xml:"short_address,omitempty"`
	// 租赁结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 操作类型
	OpType string `json:"op_type,omitempty" xml:"op_type,omitempty"`
	// 主图地址
	Pic string `json:"pic,omitempty" xml:"pic,omitempty"`
	// 经销商编码
	AgencyCode string `json:"agency_code,omitempty" xml:"agency_code,omitempty"`
	// 营业时间
	BusinessTime string `json:"business_time,omitempty" xml:"business_time,omitempty"`
	// 经营品牌编号
	BusinessBrandids string `json:"business_brandids,omitempty" xml:"business_brandids,omitempty"`
	// 城市名
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 摊位部门Id
	StallDeptid string `json:"stall_deptid,omitempty" xml:"stall_deptid,omitempty"`
	// 卖场编码
	MallCode string `json:"mall_code,omitempty" xml:"mall_code,omitempty"`
	// 品牌系列
	BrandSeries string `json:"brand_series,omitempty" xml:"brand_series,omitempty"`
	// 幂等id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 纬度
	Lat string `json:"lat,omitempty" xml:"lat,omitempty"`
	// 租赁开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 经度
	Lng string `json:"lng,omitempty" xml:"lng,omitempty"`
	// 区名称
	DistrictName string `json:"district_name,omitempty" xml:"district_name,omitempty"`
	// 联系电话
	Contract string `json:"contract,omitempty" xml:"contract,omitempty"`
	// 经营品牌名称
	BusinessBrandnames string `json:"business_brandnames,omitempty" xml:"business_brandnames,omitempty"`
	// 摊位名称
	StallName string `json:"stall_name,omitempty" xml:"stall_name,omitempty"`
	// 主营品类
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 省名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 主营品牌Id
	BrandId string `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 主营品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 卖场部门Id
	MallDepid string `json:"mall_depid,omitempty" xml:"mall_depid,omitempty"`
	// 关联的编码
	RelateCode string `json:"relate_code,omitempty" xml:"relate_code,omitempty"`
	// 门店名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 区编码
	DistrictCode int64 `json:"district_code,omitempty" xml:"district_code,omitempty"`
	// 市编码
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 省编码
	ProvinceCode int64 `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 摊位状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 摊位Id
	StallId int64 `json:"stall_id,omitempty" xml:"stall_id,omitempty"`
	// 门店id1
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
}
