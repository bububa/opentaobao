package ascpchannel

// WaybillQueryResponseData 结构体
type WaybillQueryResponseData struct {
	// 电子面单信息
	WaybillCloudPrintDtoList []WaybillCloudPrintDto `json:"waybill_cloud_print_dto_list,omitempty" xml:"waybill_cloud_print_dto_list>waybill_cloud_print_dto,omitempty"`
	// 配资源code
	CpResCode string `json:"cp_res_code,omitempty" xml:"cp_res_code,omitempty"`
	// 配资源名称
	CpResName string `json:"cp_res_name,omitempty" xml:"cp_res_name,omitempty"`
	// 配资源品牌code
	CpBrandCode string `json:"cp_brand_code,omitempty" xml:"cp_brand_code,omitempty"`
	// 物流公司名称
	LogisticsCompany string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
}
