package tmallsc

// ApplicationInfoDto 结构体
type ApplicationInfoDto struct {
	// 品牌授权商openid
	BrandLicensorOpenId string `json:"brand_licensor_open_id,omitempty" xml:"brand_licensor_open_id,omitempty"`
	// 服务商openid
	SupplierOpenId string `json:"supplier_open_id,omitempty" xml:"supplier_open_id,omitempty"`
	// 备件出库时间
	SparePartsOutTime string `json:"spare_parts_out_time,omitempty" xml:"spare_parts_out_time,omitempty"`
	// 备件申请单id
	SparePartsApplicationId int64 `json:"spare_parts_application_id,omitempty" xml:"spare_parts_application_id,omitempty"`
}
