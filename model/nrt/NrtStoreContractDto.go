package nrt

// NrtStoreContractDto 结构体
type NrtStoreContractDto struct {
	// 业务身份
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 详细地址
	ShortAddress string `json:"short_address,omitempty" xml:"short_address,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 街道名称
	TownName string `json:"town_name,omitempty" xml:"town_name,omitempty"`
	// 省份名称
	ProvinceName string `json:"province_name,omitempty" xml:"province_name,omitempty"`
	// 摊位编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 卖场名称
	MallName string `json:"mall_name,omitempty" xml:"mall_name,omitempty"`
	// 老摊位编码
	OldStoreContractCode string `json:"old_store_contract_code,omitempty" xml:"old_store_contract_code,omitempty"`
	// 合同号
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 经销商编码
	AgencyCode string `json:"agency_code,omitempty" xml:"agency_code,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 城市名称
	CityName string `json:"city_name,omitempty" xml:"city_name,omitempty"`
	// 卖场编码
	MallCode string `json:"mall_code,omitempty" xml:"mall_code,omitempty"`
	// 摊位名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 合同开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 幂等id
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 区名称
	DistrictName string `json:"district_name,omitempty" xml:"district_name,omitempty"`
	// 经销商名称
	AgencyName string `json:"agency_name,omitempty" xml:"agency_name,omitempty"`
	// 老合同号
	OldCode string `json:"old_code,omitempty" xml:"old_code,omitempty"`
	// 合同结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 联系电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 操作类型
	OpType string `json:"op_type,omitempty" xml:"op_type,omitempty"`
	// 街道编码
	TownCode int64 `json:"town_code,omitempty" xml:"town_code,omitempty"`
	// 老摊位钉钉部门
	OldStoreDingDeptId int64 `json:"old_store_ding_dept_id,omitempty" xml:"old_store_ding_dept_id,omitempty"`
	// 经度
	Lat int64 `json:"lat,omitempty" xml:"lat,omitempty"`
	// 纬度
	Lng int64 `json:"lng,omitempty" xml:"lng,omitempty"`
	// 品牌id
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 合同状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 区编码
	DistrictCode int64 `json:"district_code,omitempty" xml:"district_code,omitempty"`
	// 市编码
	CityCode int64 `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 卖场钉钉部门id
	MallDingDeptId int64 `json:"mall_ding_dept_id,omitempty" xml:"mall_ding_dept_id,omitempty"`
	// 摊位钉钉部门id
	StoreDingDeptId int64 `json:"store_ding_dept_id,omitempty" xml:"store_ding_dept_id,omitempty"`
	// 省编码
	ProvinceCode int64 `json:"province_code,omitempty" xml:"province_code,omitempty"`
}
