package btrip

// OpenVehicleOrderRs 结构体
type OpenVehicleOrderRs struct {
	// 用车原因：TRAVEL: 差旅, TRAFFIC: 市内交通, WORK: 加班, OTHER: 其它
	BusinessCategory string `json:"business_category,omitempty" xml:"business_category,omitempty"`
	// 车牌号
	CarLicense string `json:"car_license,omitempty" xml:"car_license,omitempty"`
	// 打车服务类型 1:出租车, 2:专车, 3:快车
	ServiceType int64 `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 行驶公里数
	TravelDistance string `json:"travel_distance,omitempty" xml:"travel_distance,omitempty"`
	// 取消时间
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// 司机到达目的地时间
	DriverConfirmTime string `json:"driver_confirm_time,omitempty" xml:"driver_confirm_time,omitempty"`
	// 乘客上车时间
	TakenTime string `json:"taken_time,omitempty" xml:"taken_time,omitempty"`
	// 乘客发布用车时间
	PublishTime string `json:"publish_time,omitempty" xml:"publish_time,omitempty"`
	// 订单预估价格
	EstimatePrice string `json:"estimate_price,omitempty" xml:"estimate_price,omitempty"`
	// 车辆类型
	CarInfo string `json:"car_info,omitempty" xml:"car_info,omitempty"`
	// 类型级别：1经济型、2舒适型、3豪华型
	CarLevel int64 `json:"car_level,omitempty" xml:"car_level,omitempty"`
	// 0:初始状态 1:创建失败 2: 派单成功 3：派单失败 4: 已退款 5: 已支付 6: 取消成功 7: 冻结账户失败 8： 已超时
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 打车事由
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 目的城市
	ToCityName string `json:"to_city_name,omitempty" xml:"to_city_name,omitempty"`
	// 出发城市
	FromCityName string `json:"from_city_name,omitempty" xml:"from_city_name,omitempty"`
	// 目的地
	ToAddress string `json:"to_address,omitempty" xml:"to_address,omitempty"`
	// 出发地
	FromAddress string `json:"from_address,omitempty" xml:"from_address,omitempty"`
	// 商旅系统审批单id
	ApplyId int64 `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 部门名称
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 部门id
	DepartId string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 预定人id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 预定人名称
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 企业名称
	CorpName string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 乘客名称
	PassengerName string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// 成本中心名称
	CostCenterName string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	// 商旅成本中心id
	CostCenterId int64 `json:"cost_center_id,omitempty" xml:"cost_center_id,omitempty"`
	// 实际到达城市
	RealToCityName string `json:"real_to_city_name,omitempty" xml:"real_to_city_name,omitempty"`
	// 实际出发城市
	RealFromCityName string `json:"real_from_city_name,omitempty" xml:"real_from_city_name,omitempty"`
	// 价目详情列表
	PriceInfoList []OpenPriceInfo `json:"price_info_list,omitempty" xml:"price_info_list>open_price_info,omitempty"`
	// 项目名称
	ProjectTitle string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// 项目标号
	ProjectCode string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	// 发票抬头
	InvoiceTitle string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 商旅发票id
	InvoiceId int64 `json:"invoice_id,omitempty" xml:"invoice_id,omitempty"`
	// 成本中心编号
	CostCenterNumber string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// 商旅审批单展示id(非审批api对接企业)
	ApplyShowId string `json:"apply_show_id,omitempty" xml:"apply_show_id,omitempty"`
	// 订单id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 订单更新时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 第三方行程id
	ThirdpartItineraryId string `json:"thirdpart_itinerary_id,omitempty" xml:"thirdpart_itinerary_id,omitempty"`
	// 出行人信息
	UserAffiliateList []OpenUserAffiliateDo `json:"user_affiliate_list,omitempty" xml:"user_affiliate_list>open_user_affiliate_do,omitempty"`
}
