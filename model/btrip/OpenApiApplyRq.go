package btrip

// OpenApiApplyRq 结构体
type OpenApiApplyRq struct {
	// 行程列表
	ItineraryList []OpenItineraryInfo `json:"itinerary_list,omitempty" xml:"itinerary_list>open_itinerary_info,omitempty"`
	// 出行人列表
	TravelerList []OpenUserInfo `json:"traveler_list,omitempty" xml:"traveler_list>open_user_info,omitempty"`
	// 外部出行人列表
	ExternalTravelerList []OpenUserInfo `json:"external_traveler_list,omitempty" xml:"external_traveler_list>open_user_info,omitempty"`
	// 内部人员差标列表
	TravelerStandard []InternalUserStandard `json:"traveler_standard,omitempty" xml:"traveler_standard>internal_user_standard,omitempty"`
	// 城市集行程列表
	ItinerarySetList []OpenItineraryInfo `json:"itinerary_set_list,omitempty" xml:"itinerary_set_list>open_itinerary_info,omitempty"`
	// 外部申请单id
	ThirdpartApplyId string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// 用户展示的外部审批单id信息
	ThirdpartBusinessId string `json:"thirdpart_business_id,omitempty" xml:"thirdpart_business_id,omitempty"`
	// 第三方企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 企业名称
	CorpName string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// 部门id，如果不传，会根据user相关信息去获取对应的部门信息，如果传的是错误的部门信息，后面无法做部门的费用归属；部门ID只能是数字
	DepartId string `json:"depart_id,omitempty" xml:"depart_id,omitempty"`
	// 部门名称
	DepartName string `json:"depart_name,omitempty" xml:"depart_name,omitempty"`
	// 出差事由
	TripCause string `json:"trip_cause,omitempty" xml:"trip_cause,omitempty"`
	// 申请单标题
	TripTitle string `json:"trip_title,omitempty" xml:"trip_title,omitempty"`
	// 申请人Id（第三方用户id）
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 用户名称，如果要传必须传真实姓名，如果不传则会以系统当前维护userId对应的名称进行预订
	UserName string `json:"user_name,omitempty" xml:"user_name,omitempty"`
	// 关联单号
	UnionNo string `json:"union_no,omitempty" xml:"union_no,omitempty"`
	// 审批单状态，不传入默认为0：0审批中，1同意，2拒绝
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 出差天数
	TripDay int64 `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	// 1 老版本 2 isv对外版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 1：代提交 2：本人提交 注意：当申请单为代提交时，申请单提交人自己无法为自己下单
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 审批单酒店预算，单位分。注意：特殊授权人住店产生的杂费（用餐、房间内商品等）将在退房时扣款，由企业支付
	HotelBudget int64 `json:"hotel_budget,omitempty" xml:"hotel_budget,omitempty"`
	// 审批单机票预算，单位分
	FlightBudget int64 `json:"flight_budget,omitempty" xml:"flight_budget,omitempty"`
	// 审批单火车票预算，单位分。注意：例如坐席同为”硬卧“，上/中/下铺存在价格差异的情况，预订时将按照其中最高价格，校验预算余额
	TrainBudget int64 `json:"train_budget,omitempty" xml:"train_budget,omitempty"`
	// 审批单用车预算，单位分。注意：打车场景存在不可控因素，会超出原预估价格：1. 乘客线下修改目的地；2. 堵车等道路意外情况；3. 司机添加附加费，如过路费、高速费、等待费等
	VehicleBudget int64 `json:"vehicle_budget,omitempty" xml:"vehicle_budget,omitempty"`
	// 审批单总预算，单位分
	Budget int64 `json:"budget,omitempty" xml:"budget,omitempty"`
	// 多个申请单预算合并。1：否，【union_no】相同的【申请单(apply id)】，每个的【预算】仅对本申请单生效。2：是，所有【union_no】相同的【申请单(apply id)】，其中全部【预算】合并求和，可以混用。
	BudgetMerge int64 `json:"budget_merge,omitempty" xml:"budget_merge,omitempty"`
	// 0：不限制出行人，1：限申请单内的出行人。注意：不限出行人，实际出行人也不限制差标，而且传入的出行人信息也不会存储
	LimitTraveler int64 `json:"limit_traveler,omitempty" xml:"limit_traveler,omitempty"`
	// 同时预订(机票&amp;火车票)规则。1：就高；2：就低。
	TogetherBookRule int64 `json:"together_book_rule,omitempty" xml:"together_book_rule,omitempty"`
	// 酒店合住规则
	HotelShare *HotelShareInfo `json:"hotel_share,omitempty" xml:"hotel_share,omitempty"`
	// 外部出行人差标
	ExternalTravelerStandard *ExternalUserStandard `json:"external_traveler_standard,omitempty" xml:"external_traveler_standard,omitempty"`
	// 申请单城市规则： 0出发&amp;目的地一对一，按列表传行程  1多选N个地点，城市集行程 不传默认为0 会根据商旅管理后台-通用差旅设置-行程城市规则中的设置，校验申请单本字段的值是否正确 当行程城市规则中设置的是“1对1行程”时，必须传0 当行程城市规则中设置的是“多对多城市集行程”时，必须传1 会根据此字段传入的值，校验行程传参是否正确 当申请单城市规则为0，itinerary_list行程列表必填 当申请单城市规则为1，城市集行程必填
	ItineraryRule int64 `json:"itinerary_rule,omitempty" xml:"itinerary_rule,omitempty"`
}
