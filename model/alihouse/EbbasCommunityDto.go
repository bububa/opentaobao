package alihouse

// EbbasCommunityDto 结构体
type EbbasCommunityDto struct {
	// 小区别名
	AliasNames []string `json:"alias_names,omitempty" xml:"alias_names>string,omitempty"`
	// 小区标签
	TagCodes []string `json:"tag_codes,omitempty" xml:"tag_codes>string,omitempty"`
	// 物业类型(1-住宅 2-商业)
	EstateTypes []string `json:"estate_types,omitempty" xml:"estate_types>string,omitempty"`
	// 产权年限码
	PropertyRightsYearsCodes []string `json:"property_rights_years_codes,omitempty" xml:"property_rights_years_codes>string,omitempty"`
	// 建筑类别
	BuildingCategorys []string `json:"building_categorys,omitempty" xml:"building_categorys>string,omitempty"`
	// 建筑类型
	BuildingTypes []string `json:"building_types,omitempty" xml:"building_types>string,omitempty"`
	// 房屋类型
	HouseTypes []string `json:"house_types,omitempty" xml:"house_types>string,omitempty"`
	// 户型列表
	Layouts []SyncProjectLayoutDto `json:"layouts,omitempty" xml:"layouts>sync_project_layout_dto,omitempty"`
	// 图片列表
	Pictures []ProjectPictureDto `json:"pictures,omitempty" xml:"pictures>project_picture_dto,omitempty"`
	// 装修类型，参见枚举
	DecorationStandardCodes []string `json:"decoration_standard_codes,omitempty" xml:"decoration_standard_codes>string,omitempty"`
	// 小区e码
	ECode string `json:"e_code,omitempty" xml:"e_code,omitempty"`
	// 外部id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 小区名称
	ProjectName string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// 小区亮点
	Highlights string `json:"highlights,omitempty" xml:"highlights,omitempty"`
	// 小区均价(元/m2)
	AvgPrice string `json:"avg_price,omitempty" xml:"avg_price,omitempty"`
	// 小区展示价（元）
	ShowPrice string `json:"show_price,omitempty" xml:"show_price,omitempty"`
	// 建筑面积(平米)
	BuildingArea string `json:"building_area,omitempty" xml:"building_area,omitempty"`
	// 小区详细地址
	ProjectDetailAddress string `json:"project_detail_address,omitempty" xml:"project_detail_address,omitempty"`
	// 地图定位-详细地址
	MapLocationDetailAddress string `json:"map_location_detail_address,omitempty" xml:"map_location_detail_address,omitempty"`
	// 地图定位-经度
	MapLocationLongitude string `json:"map_location_longitude,omitempty" xml:"map_location_longitude,omitempty"`
	// 地图定位-纬度
	MapLocationLatitude string `json:"map_location_latitude,omitempty" xml:"map_location_latitude,omitempty"`
	// 开发商全称
	DeveloperFullName string `json:"developer_full_name,omitempty" xml:"developer_full_name,omitempty"`
	// 小区概况-占地面积(平米)
	CommunityArea string `json:"community_area,omitempty" xml:"community_area,omitempty"`
	// 小区概况-容积率(例：20)（单位：%）
	CommunityVolumeRate string `json:"community_volume_rate,omitempty" xml:"community_volume_rate,omitempty"`
	// 小区概况-绿化率(例：20)（单位：%）
	CommunityGreeningRate string `json:"community_greening_rate,omitempty" xml:"community_greening_rate,omitempty"`
	// 车位比例
	ParkingRadio string `json:"parking_radio,omitempty" xml:"parking_radio,omitempty"`
	// 小区概况-物业公司
	CommunityEstateCompany string `json:"community_estate_company,omitempty" xml:"community_estate_company,omitempty"`
	// 小区概况-物业费用(元/m?/月)
	CommunityEstateExpenses string `json:"community_estate_expenses,omitempty" xml:"community_estate_expenses,omitempty"`
	// 供暖方式描述
	CommunityHeatingTypeDesc string `json:"community_heating_type_desc,omitempty" xml:"community_heating_type_desc,omitempty"`
	// 供水方式描述
	CommunityWaterTypeDesc string `json:"community_water_type_desc,omitempty" xml:"community_water_type_desc,omitempty"`
	// 供电方式描述
	CommunityPowerTypeDesc string `json:"community_power_type_desc,omitempty" xml:"community_power_type_desc,omitempty"`
	// 小区介绍
	OverallIntroduction string `json:"overall_introduction,omitempty" xml:"overall_introduction,omitempty"`
	// 此次修改的原因
	SubmitReason string `json:"submit_reason,omitempty" xml:"submit_reason,omitempty"`
	// 小区所属国家名字
	CountryName string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	// 小区所属省名字
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 小区所属市名字
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 小区所属区名字
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 小区所属街道名字
	StreetName string `json:"street_name,omitempty" xml:"street_name,omitempty"`
	// 环线
	LoopLine string `json:"loop_line,omitempty" xml:"loop_line,omitempty"`
	// 车位费(单位：元/月)
	ParkingFee string `json:"parking_fee,omitempty" xml:"parking_fee,omitempty"`
	// 供气描述
	GasSupplyDescription string `json:"gas_supply_description,omitempty" xml:"gas_supply_description,omitempty"`
	// 主力户型描述
	MainHouseTypeDescription string `json:"main_house_type_description,omitempty" xml:"main_house_type_description,omitempty"`
	// 是否人车分流
	DiversionPeopleVehicles string `json:"diversion_people_vehicles,omitempty" xml:"diversion_people_vehicles,omitempty"`
	// 围栏数据
	Fencing string `json:"fencing,omitempty" xml:"fencing,omitempty"`
	// 建筑年代
	ArchitectureAge string `json:"architecture_age,omitempty" xml:"architecture_age,omitempty"`
	// 物业服务中心地址
	PropertyAddress string `json:"property_address,omitempty" xml:"property_address,omitempty"`
	// 物业工作时间
	PropertyWorkingHours string `json:"property_working_hours,omitempty" xml:"property_working_hours,omitempty"`
	// 物业电话
	PropertyTelephone string `json:"property_telephone,omitempty" xml:"property_telephone,omitempty"`
	// 小区是否封闭
	IsCommunityClosed string `json:"is_community_closed,omitempty" xml:"is_community_closed,omitempty"`
	// 固定停车费标准
	FixedParkingFeeStandard string `json:"fixed_parking_fee_standard,omitempty" xml:"fixed_parking_fee_standard,omitempty"`
	// 临时停车费标准
	TempParkingFeeStandard string `json:"temp_parking_fee_standard,omitempty" xml:"temp_parking_fee_standard,omitempty"`
	// 外立面风格
	FacadeStyle string `json:"facade_style,omitempty" xml:"facade_style,omitempty"`
	// 外墙材料
	ExteriorWallMaterials string `json:"exterior_wall_materials,omitempty" xml:"exterior_wall_materials,omitempty"`
	// 是否安装智能道闸
	IsIntelligentGateInstalled string `json:"is_intelligent_gate_installed,omitempty" xml:"is_intelligent_gate_installed,omitempty"`
	// 是否有门禁
	HasAccessControl string `json:"has_access_control,omitempty" xml:"has_access_control,omitempty"`
	// 是否有监控
	HasMonitor string `json:"has_monitor,omitempty" xml:"has_monitor,omitempty"`
	// 是否24小时值班
	IsDutyAllDay string `json:"is_duty_all_day,omitempty" xml:"is_duty_all_day,omitempty"`
	// 保安巡逻频次
	PatrolFrequency string `json:"patrol_frequency,omitempty" xml:"patrol_frequency,omitempty"`
	// 是否110联网
	HasPoliceConnected string `json:"has_police_connected,omitempty" xml:"has_police_connected,omitempty"`
	// 小区等级
	CommunityLevel string `json:"community_level,omitempty" xml:"community_level,omitempty"`
	// 小区拼音首字母
	Pinyin string `json:"pinyin,omitempty" xml:"pinyin,omitempty"`
	// 小区名拼音
	FullPinyin string `json:"full_pinyin,omitempty" xml:"full_pinyin,omitempty"`
	// 小区内部配套
	InternalMatching string `json:"internal_matching,omitempty" xml:"internal_matching,omitempty"`
	// 商圈ID
	BuisnessId string `json:"buisness_id,omitempty" xml:"buisness_id,omitempty"`
	// 板块ID
	ModuleId string `json:"module_id,omitempty" xml:"module_id,omitempty"`
	// 均价单位
	AvgPriceUnit string `json:"avg_price_unit,omitempty" xml:"avg_price_unit,omitempty"`
	// 混合功能区ID
	FunctionAreaId string `json:"function_area_id,omitempty" xml:"function_area_id,omitempty"`
	// 地铁
	MetroLine string `json:"metro_line,omitempty" xml:"metro_line,omitempty"`
	// 功能区ID
	FunctionId string `json:"function_id,omitempty" xml:"function_id,omitempty"`
	// 高德小区名称
	AmapProjectName string `json:"amap_project_name,omitempty" xml:"amap_project_name,omitempty"`
	// 高德定位-详细地址
	AmapLocationDetailAddress string `json:"amap_location_detail_address,omitempty" xml:"amap_location_detail_address,omitempty"`
	// 高德小区经度
	AmapLocationLongitude string `json:"amap_location_longitude,omitempty" xml:"amap_location_longitude,omitempty"`
	// 高德小区纬度
	AmapLocationLatitude string `json:"amap_location_latitude,omitempty" xml:"amap_location_latitude,omitempty"`
	// 来源
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 小区概况-楼栋数
	CommunityPlanningBuilding int64 `json:"community_planning_building,omitempty" xml:"community_planning_building,omitempty"`
	// 小区概况-总户数
	CommunityPlanningHouseholds int64 `json:"community_planning_households,omitempty" xml:"community_planning_households,omitempty"`
	// 小区所属国家Id
	CountryId int64 `json:"country_id,omitempty" xml:"country_id,omitempty"`
	// 小区所属省编码
	ProvId int64 `json:"prov_id,omitempty" xml:"prov_id,omitempty"`
	// 小区所属市编码
	CityId int64 `json:"city_id,omitempty" xml:"city_id,omitempty"`
	// 小区所属区编码
	AreaId int64 `json:"area_id,omitempty" xml:"area_id,omitempty"`
	// 小区所属街道编码
	StreetId int64 `json:"street_id,omitempty" xml:"street_id,omitempty"`
	// 车位数
	ParkingNumber int64 `json:"parking_number,omitempty" xml:"parking_number,omitempty"`
	// 地上车位数
	OvergroundParkingNumber int64 `json:"overground_parking_number,omitempty" xml:"overground_parking_number,omitempty"`
	// 地下车位数
	UndergroundParkingNumber int64 `json:"underground_parking_number,omitempty" xml:"underground_parking_number,omitempty"`
	// 交易权属
	TransactionOwnership int64 `json:"transaction_ownership,omitempty" xml:"transaction_ownership,omitempty"`
	// 非机动车库标识
	IdentificationNonMotorizedGarage int64 `json:"identification_non_motorized_garage,omitempty" xml:"identification_non_motorized_garage,omitempty"`
	// 空调新风标识
	AirFreshAirSign int64 `json:"air_fresh_air_sign,omitempty" xml:"air_fresh_air_sign,omitempty"`
	// 保安岗亭数
	SecurityPostsNumber int64 `json:"security_posts_number,omitempty" xml:"security_posts_number,omitempty"`
	// 保安人数
	SecurityPersonnelNumber int64 `json:"security_personnel_number,omitempty" xml:"security_personnel_number,omitempty"`
	// 车位配置码
	GarageConfiguration int64 `json:"garage_configuration,omitempty" xml:"garage_configuration,omitempty"`
	// 总楼层
	TotalFloor int64 `json:"total_floor,omitempty" xml:"total_floor,omitempty"`
	// 是否有电梯
	HasElevator int64 `json:"has_elevator,omitempty" xml:"has_elevator,omitempty"`
	// 公寓管理信息对象
	ApartmentManagement *ApartmentManagementDto `json:"apartment_management,omitempty" xml:"apartment_management,omitempty"`
	// 城市ID-临时
	TempCityId int64 `json:"temp_city_id,omitempty" xml:"temp_city_id,omitempty"`
	// 区域ID-临时
	TempAreaId int64 `json:"temp_area_id,omitempty" xml:"temp_area_id,omitempty"`
	// 功能区ID-临时
	TempFunctionId int64 `json:"temp_function_id,omitempty" xml:"temp_function_id,omitempty"`
	// 板块ID-临时
	TempModuleId int64 `json:"temp_module_id,omitempty" xml:"temp_module_id,omitempty"`
	// 是否为测试数据true:是
	IsTest bool `json:"is_test,omitempty" xml:"is_test,omitempty"`
}
