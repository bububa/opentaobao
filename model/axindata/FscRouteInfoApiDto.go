package axindata

// FscRouteInfoApiDto 结构体
type FscRouteInfoApiDto struct {
	// 多媒体数据列表
	MediaList []FscProductMediaApiDto `json:"media_list,omitempty" xml:"media_list>fsc_product_media_api_dto,omitempty"`
	// 线路主题ID列表
	RouteLabelIdList []string `json:"route_label_id_list,omitempty" xml:"route_label_id_list>string,omitempty"`
	// 出行人定义
	TravellerDefineList []FscTravellerDefineApiDto `json:"traveller_define_list,omitempty" xml:"traveller_define_list>fsc_traveller_define_api_dto,omitempty"`
	// POI的Id列表
	PoiIdList []string `json:"poi_id_list,omitempty" xml:"poi_id_list>string,omitempty"`
	// 建议售卖渠道列表
	SalesInfoList []FscSalesInfoApiDto `json:"sales_info_list,omitempty" xml:"sales_info_list>fsc_sales_info_api_dto,omitempty"`
	// 产品推荐特色（最多3条）
	Recommends []string `json:"recommends,omitempty" xml:"recommends>string,omitempty"`
	// 线路行程列表
	JourneyList []FscRouteJourneyApiDto `json:"journey_list,omitempty" xml:"journey_list>fsc_route_journey_api_dto,omitempty"`
	// 附加规则列表
	ProductRules []FscProductRuleApiDto `json:"product_rules,omitempty" xml:"product_rules>fsc_product_rule_api_dto,omitempty"`
	// 取消政策列表
	CancelPolicyList []FscProductCancelPolicyApiDto `json:"cancel_policy_list,omitempty" xml:"cancel_policy_list>fsc_product_cancel_policy_api_dto,omitempty"`
	// 线路编号
	RouteCode string `json:"route_code,omitempty" xml:"route_code,omitempty"`
	// 线路标题
	RouteName string `json:"route_name,omitempty" xml:"route_name,omitempty"`
	// 线路副标题
	RouteSubName string `json:"route_sub_name,omitempty" xml:"route_sub_name,omitempty"`
	// 产品类别 INTL_GROUP_TRAVEL：出境跟团游 DOM_GROUP_TRAVEL：境内跟团游
	SubCategory string `json:"sub_category,omitempty" xml:"sub_category,omitempty"`
	// 业务区域编码
	BusinessAreaId string `json:"business_area_id,omitempty" xml:"business_area_id,omitempty"`
	// 出发城市Id
	StartCityId string `json:"start_city_id,omitempty" xml:"start_city_id,omitempty"`
	// 目的城市Id
	EndCityId string `json:"end_city_id,omitempty" xml:"end_city_id,omitempty"`
	// 出发国家Id
	StartCountryId string `json:"start_country_id,omitempty" xml:"start_country_id,omitempty"`
	// 目的国家Id
	EndCountryId string `json:"end_country_id,omitempty" xml:"end_country_id,omitempty"`
	// 线路主图
	RouteMainPic string `json:"route_main_pic,omitempty" xml:"route_main_pic,omitempty"`
	// 线路特色
	RouteFeature string `json:"route_feature,omitempty" xml:"route_feature,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 线路产品备注
	RouteRemark string `json:"route_remark,omitempty" xml:"route_remark,omitempty"`
	// 确认规则 INSTANT_CONFIRMATION，即时确认 SECOND_CONFIRMATION，二次确认
	ConfirmType string `json:"confirm_type,omitempty" xml:"confirm_type,omitempty"`
	// 线路状态 ON_SHELF：在架 DOWN_SHELF：下架
	RouteStatus string `json:"route_status,omitempty" xml:"route_status,omitempty"`
	// 签证信息
	VisaName string `json:"visa_name,omitempty" xml:"visa_name,omitempty"`
	// 联系人姓名
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人电话
	ContactPhone string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	// 供应商底层线路编码
	SupplierRouteCode string `json:"supplier_route_code,omitempty" xml:"supplier_route_code,omitempty"`
	// 参团类型 1:出发地参团 2:目的地参团
	GroupType int64 `json:"group_type,omitempty" xml:"group_type,omitempty"`
	// 行程天数
	RouteDay int64 `json:"route_day,omitempty" xml:"route_day,omitempty"`
	// 行程晚数
	RouteNight int64 `json:"route_night,omitempty" xml:"route_night,omitempty"`
	// 是否含保险
	IncludeInsuranceFlag bool `json:"include_insurance_flag,omitempty" xml:"include_insurance_flag,omitempty"`
}
