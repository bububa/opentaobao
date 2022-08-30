package alihouse

// EbbasProjectDto 结构体
type EbbasProjectDto struct {
	// 楼盘别名
	AliasNames []string `json:"alias_names,omitempty" xml:"alias_names>string,omitempty"`
	// 标签码
	TagCodes []string `json:"tag_codes,omitempty" xml:"tag_codes>string,omitempty"`
	// 物业类型,多选
	EstateTypes []string `json:"estate_types,omitempty" xml:"estate_types>string,omitempty"`
	// 产权年限码
	PropertyRightsYearsCodes []string `json:"property_rights_years_codes,omitempty" xml:"property_rights_years_codes>string,omitempty"`
	// 装修类型，参见枚举
	DecorationStandardCodes []string `json:"decoration_standard_codes,omitempty" xml:"decoration_standard_codes>string,omitempty"`
	// 建筑类别
	BuildingCategorys []string `json:"building_categorys,omitempty" xml:"building_categorys>string,omitempty"`
	// 建筑类型
	BuildingTypes []string `json:"building_types,omitempty" xml:"building_types>string,omitempty"`
	// 房屋类型
	HouseTypes []string `json:"house_types,omitempty" xml:"house_types>string,omitempty"`
	// 周边商业
	PeripheralBusiness []string `json:"peripheral_business,omitempty" xml:"peripheral_business>string,omitempty"`
	// 周边景观
	SurroundingLandscape []string `json:"surrounding_landscape,omitempty" xml:"surrounding_landscape>string,omitempty"`
	// 周边公园
	SurroundingParks []string `json:"surrounding_parks,omitempty" xml:"surrounding_parks>string,omitempty"`
	// 周边医院
	SurroundingHospitals []string `json:"surrounding_hospitals,omitempty" xml:"surrounding_hospitals>string,omitempty"`
	// 周边学校
	SurroundingSchools []string `json:"surrounding_schools,omitempty" xml:"surrounding_schools>string,omitempty"`
	// 周边交通
	SurroundingTraffic []string `json:"surrounding_traffic,omitempty" xml:"surrounding_traffic>string,omitempty"`
	// 周边餐饮
	SurroundingRestaurant []string `json:"surrounding_restaurant,omitempty" xml:"surrounding_restaurant>string,omitempty"`
	// 周边银行
	SurroundingBanks []string `json:"surrounding_banks,omitempty" xml:"surrounding_banks>string,omitempty"`
	// 投资商
	Investors []string `json:"investors,omitempty" xml:"investors>string,omitempty"`
	// 户型对象列表
	Layouts []SyncProjectLayoutDto `json:"layouts,omitempty" xml:"layouts>sync_project_layout_dto,omitempty"`
	// 图片对象列表
	Pictures []ProjectPictureDto `json:"pictures,omitempty" xml:"pictures>project_picture_dto,omitempty"`
	// 楼盘ecode
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 楼盘名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 备案名
	RecordName string `json:"record_name,omitempty" xml:"record_name,omitempty"`
	// 楼盘亮点
	Highlights string `json:"highlights,omitempty" xml:"highlights,omitempty"`
	// 楼盘总价(万元/套)
	TotalPrice string `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 楼盘总价单位
	TotalPriceUnit string `json:"total_price_unit,omitempty" xml:"total_price_unit,omitempty"`
	// 楼盘均价(元/㎡)
	AvgPrice string `json:"avg_price,omitempty" xml:"avg_price,omitempty"`
	// 楼盘均价单位
	AvgPriceUnit string `json:"avg_price_unit,omitempty" xml:"avg_price_unit,omitempty"`
	// 楼盘展示价（元）
	ShowPrice string `json:"show_price,omitempty" xml:"show_price,omitempty"`
	// 平均建筑面积
	BuildingArea string `json:"building_area,omitempty" xml:"building_area,omitempty"`
	// 楼盘详细地址
	ProjectDetailAddress string `json:"project_detail_address,omitempty" xml:"project_detail_address,omitempty"`
	// 地图定位-详细地址
	MapLocationDetailAddress string `json:"map_location_detail_address,omitempty" xml:"map_location_detail_address,omitempty"`
	// 地图定位-经度
	MapLocationLongitude string `json:"map_location_longitude,omitempty" xml:"map_location_longitude,omitempty"`
	// 地图定位-纬度
	MapLocationLatitude string `json:"map_location_latitude,omitempty" xml:"map_location_latitude,omitempty"`
	// 开发商全称
	DeveloperFullName string `json:"developer_full_name,omitempty" xml:"developer_full_name,omitempty"`
	// 品牌商名称
	BrandName string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// 售楼处地址
	DeveloperSalesOfficeAddress string `json:"developer_sales_office_address,omitempty" xml:"developer_sales_office_address,omitempty"`
	// 开盘时间(yyyy-MM)
	DeveloperOpeningTime string `json:"developer_opening_time,omitempty" xml:"developer_opening_time,omitempty"`
	// 交付时间(yyyy-MM)
	DeveloperDueTime string `json:"developer_due_time,omitempty" xml:"developer_due_time,omitempty"`
	// 联系人座机
	PhoneNo string `json:"phone_no,omitempty" xml:"phone_no,omitempty"`
	// 联系人手机
	Telephone string `json:"telephone,omitempty" xml:"telephone,omitempty"`
	// 400主机电话
	MainPhone string `json:"main_phone,omitempty" xml:"main_phone,omitempty"`
	// 400分机电话
	SubPhone string `json:"sub_phone,omitempty" xml:"sub_phone,omitempty"`
	// 小区概况-占地面积(平米)
	CommunityArea string `json:"community_area,omitempty" xml:"community_area,omitempty"`
	// 小区概况-容积率
	CommunityVolumeRate string `json:"community_volume_rate,omitempty" xml:"community_volume_rate,omitempty"`
	// 小区概况-绿化率
	CommunityGreeningRate string `json:"community_greening_rate,omitempty" xml:"community_greening_rate,omitempty"`
	// 车位比例
	ParkingRadio string `json:"parking_radio,omitempty" xml:"parking_radio,omitempty"`
	// 小区概况-物业公司
	CommunityEstateCompany string `json:"community_estate_company,omitempty" xml:"community_estate_company,omitempty"`
	// 小区概况-物业费用
	CommunityEstateExpenses string `json:"community_estate_expenses,omitempty" xml:"community_estate_expenses,omitempty"`
	// 供暖方式描述
	CommunityHeatingTypeDesc string `json:"community_heating_type_desc,omitempty" xml:"community_heating_type_desc,omitempty"`
	// 供水方式描述
	CommunityWaterTypeDesc string `json:"community_water_type_desc,omitempty" xml:"community_water_type_desc,omitempty"`
	// 供电方式描述
	CommunityPowerTypeDesc string `json:"community_power_type_desc,omitempty" xml:"community_power_type_desc,omitempty"`
	// 楼盘整体介绍
	OverallIntroduction string `json:"overall_introduction,omitempty" xml:"overall_introduction,omitempty"`
	// 此次修改的原因
	SubmitReason string `json:"submit_reason,omitempty" xml:"submit_reason,omitempty"`
	// 国家
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 省
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 街道
	StreetName string `json:"street_name,omitempty" xml:"street_name,omitempty"`
	// 环线
	LoopLine string `json:"loop_line,omitempty" xml:"loop_line,omitempty"`
	// 交通状况描述
	TrafficDescription string `json:"traffic_description,omitempty" xml:"traffic_description,omitempty"`
	// 楼层分布描述
	FloorDescription string `json:"floor_description,omitempty" xml:"floor_description,omitempty"`
	// 项目特色描述
	ProjectFeatureDescription string `json:"project_feature_description,omitempty" xml:"project_feature_description,omitempty"`
	// 工程进度描述
	ProjectScheduleDescription string `json:"project_schedule_description,omitempty" xml:"project_schedule_description,omitempty"`
	// 销售进度描述
	SalesProgressDescription string `json:"sales_progress_description,omitempty" xml:"sales_progress_description,omitempty"`
	// 开工时间
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 竣工时间
	CompletionTime string `json:"completion_time,omitempty" xml:"completion_time,omitempty"`
	// 得房率
	HouseAcquisitionRate string `json:"house_acquisition_rate,omitempty" xml:"house_acquisition_rate,omitempty"`
	// 付款方式
	PaymentMethod string `json:"payment_method,omitempty" xml:"payment_method,omitempty"`
	// 首付比例
	DownPaymentRatio string `json:"down_payment_ratio,omitempty" xml:"down_payment_ratio,omitempty"`
	// 车位费
	ParkingFee string `json:"parking_fee,omitempty" xml:"parking_fee,omitempty"`
	// 车库配置描述
	GarageConfigurationDescription string `json:"garage_configuration_description,omitempty" xml:"garage_configuration_description,omitempty"`
	// 供气描述
	GasSupplyDescription string `json:"gas_supply_description,omitempty" xml:"gas_supply_description,omitempty"`
	// 主力户型描述
	MainHouseTypeDescription string `json:"main_house_type_description,omitempty" xml:"main_house_type_description,omitempty"`
	// 人车分流
	DiversionPeopleVehicles string `json:"diversion_people_vehicles,omitempty" xml:"diversion_people_vehicles,omitempty"`
	// 电话区号
	PhoneAreaCode string `json:"phone_area_code,omitempty" xml:"phone_area_code,omitempty"`
	// 不动产权证
	RealEstateCertificate string `json:"real_estate_certificate,omitempty" xml:"real_estate_certificate,omitempty"`
	// 分期
	Stages string `json:"stages,omitempty" xml:"stages,omitempty"`
	// 围栏数据
	Fencing string `json:"fencing,omitempty" xml:"fencing,omitempty"`
	// 建筑年代
	ArchitectureAge string `json:"architecture_age,omitempty" xml:"architecture_age,omitempty"`
	// 小区拼音首字母
	Pinyin string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
	// 小区名拼音
	FullPinyin string `json:"full_pinyin,omitempty" xml:"full_pinyin,omitempty"`
	// 预售证
	PreSalePermit string `json:"pre_sale_permit,omitempty" xml:"pre_sale_permit,omitempty"`
	// 小区内部配套
	InternalMatching string `json:"internal_matching,omitempty" xml:"internal_matching,omitempty"`
	// 其他配套
	OtherFacilities string `json:"other_facilities,omitempty" xml:"other_facilities,omitempty"`
	// 商圈ID
	BuisnessId string `json:"buisness_id,omitempty" xml:"buisness_id,omitempty"`
	// 板块ID
	ModuleId string `json:"module_id,omitempty" xml:"module_id,omitempty"`
	// 营销方式
	DirectSale string `json:"direct_sale,omitempty" xml:"direct_sale,omitempty"`
	// 车位总数描述
	ParkingNumberDesc string `json:"parking_number_desc,omitempty" xml:"parking_number_desc,omitempty"`
	// 功能区id
	FunctionId string `json:"function_id,omitempty" xml:"function_id,omitempty"`
	// 外部项目店id
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 1-EBBAS 2-开发商 3-ETC 默认2
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// 1-新房楼盘 2-小区  默认1-新房
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 小区概况-规划楼栋
	CommunityPlanningBuilding int64 `json:"community_planning_building,omitempty" xml:"community_planning_building,omitempty"`
	// 小区概况-规划户数
	CommunityPlanningHouseholds int64 `json:"community_planning_households,omitempty" xml:"community_planning_households,omitempty"`
	// 国家编码
	CountryId int64 `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 省编码
	ProvId int64 `json:"prov_id,omitempty" xml:"prov_id,omitempty"`
	// 市编码
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 区编码
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 街道编码
	StreetId int64 `json:"street_id,omitempty" xml:"street_id,omitempty"`
	// 车位数
	ParkingNumber int64 `json:"parking_number,omitempty" xml:"parking_number,omitempty"`
	// 地上车位数
	OvergroundParkingNumber int64 `json:"overground_parking_number,omitempty" xml:"overground_parking_number,omitempty"`
	// 地下车位数
	UndergroundParkingNumber int64 `json:"underground_parking_number,omitempty" xml:"underground_parking_number,omitempty"`
	// 楼盘销售状态
	ProjectSalesStatus int64 `json:"project_sales_status,omitempty" xml:"project_sales_status,omitempty"`
	// 城市ID-临时
	TempCityId int64 `json:"temp_city_id,omitempty" xml:"temp_city_id,omitempty"`
	// 区域ID-临时
	TempAreaId int64 `json:"temp_area_id,omitempty" xml:"temp_area_id,omitempty"`
	// 功能区ID-临时
	TempFunctionId int64 `json:"temp_function_id,omitempty" xml:"temp_function_id,omitempty"`
	// 板块ID-临时
	TempModuleId int64 `json:"temp_module_id,omitempty" xml:"temp_module_id,omitempty"`
	// 是否为测试楼盘 true-是测试楼盘
	IsTest bool `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
