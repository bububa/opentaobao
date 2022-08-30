package nrt

// EaStoreContractDto 结构体
type EaStoreContractDto struct {
	// 合同号
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 老合同号
	OldCode string `json:"old_code,omitempty" xml:"old_code,omitempty"`
	// 老摊位号
	OldStoreContractCode string `json:"old_store_contract_code,omitempty" xml:"old_store_contract_code,omitempty"`
	// 卖场编码
	MallCode string `json:"mall_code,omitempty" xml:"mall_code,omitempty"`
	// 卖场名称
	MallName string `json:"mall_name,omitempty" xml:"mall_name,omitempty"`
	// 商户编码
	AgencyCode string `json:"agency_code,omitempty" xml:"agency_code,omitempty"`
	// 商户名称
	AgencyName string `json:"agency_name,omitempty" xml:"agency_name,omitempty"`
	// 摊位编码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 摊位名称
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 合同开始时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 合同结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 合同状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 品牌名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 类目名称
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 北京市xxx
	ShortAddress string `json:"short_address,omitempty" xml:"short_address,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 摊位id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 品牌id
	BrandId int64 `json:"brand_id,omitempty" xml:"brand_id,omitempty"`
	// 卖场钉钉部门id
	MallDingDeptId int64 `json:"mall_ding_dept_id,omitempty" xml:"mall_ding_dept_id,omitempty"`
	// 摊位钉钉部门id
	StoreDingDeptId int64 `json:"store_ding_dept_id,omitempty" xml:"store_ding_dept_id,omitempty"`
	// 摊位老钉钉部门id
	OldStoreDingDeptId int64 `json:"old_store_ding_dept_id,omitempty" xml:"old_store_ding_dept_id,omitempty"`
	// 数据状态
	RowStatus int64 `json:"row_status,omitempty" xml:"row_status,omitempty"`
	// 卖场id
	MallId int64 `json:"mall_id,omitempty" xml:"mall_id,omitempty"`
}
