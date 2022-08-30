package alitripbp

// ExamineOuterUserRequest 结构体
type ExamineOuterUserRequest struct {
	// 活动标识
	Activity string `json:"activity,omitempty" xml:"activity,omitempty"`
	// 渠道标识
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 用户标识
	OuterUserId string `json:"outer_user_id,omitempty" xml:"outer_user_id,omitempty"`
	// 平台标识
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
}
