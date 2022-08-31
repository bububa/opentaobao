package btrip

// OpenApiJumpInfoRq 结构体
type OpenApiJumpInfoRq struct {
	// 第三方企业ID
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 第三方用户ID
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 行程单号
	ItineraryId string `json:"itinerary_id,omitempty" xml:"itinerary_id,omitempty"`
	// 联系电话
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 预定跳转：1机票2火车3酒店4用车
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 跳转url：1 预定2 订单3 管理4 H5首页5 注册签约页面
	ActionType int64 `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 版本：1老版本2isv对外版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 注册签约时企业信息（注册签约时必填）
	CorpInfoRq *CorpInfoRq `json:"corp_info_rq,omitempty" xml:"corp_info_rq,omitempty"`
	// 注册签约时管理员信息（注册签约时必填）
	UserInfoRq *UserInfoRq `json:"user_info_rq,omitempty" xml:"user_info_rq,omitempty"`
}
