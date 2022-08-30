package alihouse

// ApartmentManagementDto 结构体
type ApartmentManagementDto struct {
	// 服务设施
	ServiceFacilitys []string `json:"service_facilitys,omitempty" xml:"service_facilitys>string,omitempty"`
	// 外部公寓管理信息id
	OuterManageId string `json:"outer_manage_id,omitempty" xml:"outer_manage_id,omitempty"`
	// 外部公司id
	OuterCompanyId string `json:"outer_company_id,omitempty" xml:"outer_company_id,omitempty"`
	// 外部门店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 外部管家id
	OuterBrokerId string `json:"outer_broker_id,omitempty" xml:"outer_broker_id,omitempty"`
	// 外部品牌id
	OuterBrandId string `json:"outer_brand_id,omitempty" xml:"outer_brand_id,omitempty"`
	// 核验编号
	VerificationCode string `json:"verification_code,omitempty" xml:"verification_code,omitempty"`
	// 公寓描述
	ApartmentDesc string `json:"apartment_desc,omitempty" xml:"apartment_desc,omitempty"`
	// 营业状态
	BusinessStatus int64 `json:"business_status,omitempty" xml:"business_status,omitempty"`
}
