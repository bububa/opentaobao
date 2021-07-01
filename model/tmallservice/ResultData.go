package tmallservice

// ResultData 结构体
type ResultData struct {
	// 页面大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 数据
	OnePagedDataList []OnePagedDataList `json:"one_paged_data_list,omitempty" xml:"one_paged_data_list>one_paged_data_list,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 街道地址编码
	TownCode string `json:"town_code,omitempty" xml:"town_code,omitempty"`
	// 网点名称
	ServiceStoreName string `json:"service_store_name,omitempty" xml:"service_store_name,omitempty"`
	// 网点负责人姓名
	ManagerName string `json:"manager_name,omitempty" xml:"manager_name,omitempty"`
	// 门店类型
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 品牌认证资料
	BrandCertification string `json:"brand_certification,omitempty" xml:"brand_certification,omitempty"`
	// 网点详细地址
	AddressDetail string `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	// 别名
	ServiceStoreAlias string `json:"service_store_alias,omitempty" xml:"service_store_alias,omitempty"`
	// 网点所在市
	AddressCity string `json:"address_city,omitempty" xml:"address_city,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty" xml:"longitude,omitempty"`
	// 营业执照照片
	LicensePhoto string `json:"license_photo,omitempty" xml:"license_photo,omitempty"`
	// 服务场所编码
	ServiceSiteCode string `json:"service_site_code,omitempty" xml:"service_site_code,omitempty"`
	// 网点所在区
	AddressDistrict string `json:"address_district,omitempty" xml:"address_district,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 网点所在省
	AddressProvince string `json:"address_province,omitempty" xml:"address_province,omitempty"`
	// 区地址编码
	DistrictCode string `json:"district_code,omitempty" xml:"district_code,omitempty"`
	// 业务类型
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 有效期
	GmtExpire string `json:"gmt_expire,omitempty" xml:"gmt_expire,omitempty"`
	// 城市地址编码
	CityCode string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 纬度
	Latitude string `json:"latitude,omitempty" xml:"latitude,omitempty"`
	// 网点所在街道
	AddressTown string `json:"address_town,omitempty" xml:"address_town,omitempty"`
	// 门头照片
	FrontPhoto string `json:"front_photo,omitempty" xml:"front_photo,omitempty"`
	// 网点地址编码
	AddressCode int64 `json:"address_code,omitempty" xml:"address_code,omitempty"`
	// 认证的天猫品牌id列表
	CertificatedBrandIds string `json:"certificated_brand_ids,omitempty" xml:"certificated_brand_ids,omitempty"`
	// 省地址编码
	ProvinceCode string `json:"province_code,omitempty" xml:"province_code,omitempty"`
	// 营业时间
	BusinessHours string `json:"business_hours,omitempty" xml:"business_hours,omitempty"`
	// 商户中心门店id
	PlaceStoreId string `json:"place_store_id,omitempty" xml:"place_store_id,omitempty"`
	// 网点编码
	ServiceStoreCode string `json:"service_store_code,omitempty" xml:"service_store_code,omitempty"`
	// 对外服务手机号
	ServiceMobile string `json:"service_mobile,omitempty" xml:"service_mobile,omitempty"`
	// 额外属性
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 统一社会信用代码
	SocialCreditCode string `json:"social_credit_code,omitempty" xml:"social_credit_code,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 服务产品title
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 服务产品介绍
	ServiceProductContent string `json:"service_product_content,omitempty" xml:"service_product_content,omitempty"`
	// 服务产品id
	ServiceProductId int64 `json:"service_product_id,omitempty" xml:"service_product_id,omitempty"`
	// 服务产品sku列表
	SimpleServiceSkuList []SimpleServiceSkuDtOs `json:"simple_service_sku_list,omitempty" xml:"simple_service_sku_list>simple_service_sku_dt_os,omitempty"`
	// 服务名称
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务产品状态
	ServiceProductStatus int64 `json:"service_product_status,omitempty" xml:"service_product_status,omitempty"`
	// 服务产品类型
	ServiceProductType int64 `json:"service_product_type,omitempty" xml:"service_product_type,omitempty"`
	// 上传结果是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 错误码
	FailureCode string `json:"failure_code,omitempty" xml:"failure_code,omitempty"`
	// 出错因子列表
	FailurePriceFactors string `json:"failure_price_factors,omitempty" xml:"failure_price_factors,omitempty"`
	// 错误详情
	FailureDetail string `json:"failure_detail,omitempty" xml:"failure_detail,omitempty"`
}
