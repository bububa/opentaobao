package btrip

// OpenApiApplyRq 结构体
type OpenApiApplyRq struct {
	// 行程列表
	ItineraryList []OpenItineraryInfo `json:"itinerary_list,omitempty" xml:"itinerary_list>open_itinerary_info,omitempty"`
	// 出行人列表
	TravelerList []OpenUserInfo `json:"traveler_list,omitempty" xml:"traveler_list>open_user_info,omitempty"`
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
	// 审批单状态，不传入默认为0：0审批中，1同意，2拒绝
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 出差天数
	TripDay int64 `json:"trip_day,omitempty" xml:"trip_day,omitempty"`
	// 1 老版本 2 isv对外版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 1：代提交 2：本人提交 注意：当申请单为代提交时，申请单提交人自己无法为自己下单
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}
