package alitripmerchant

// ActivityParticipateNumberVo 结构体
type ActivityParticipateNumberVo struct {
	// 活动id
	OfferId int64 `json:"offer_id,omitempty" xml:"offer_id,omitempty"`
	// 剩余次数
	LeftTime int64 `json:"left_time,omitempty" xml:"left_time,omitempty"`
}
