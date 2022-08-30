package alitripmerchant

// PointsDetailDto 结构体
type PointsDetailDto struct {
	// 会员奖励积分
	NbPoints int64 `json:"nb_points,omitempty" xml:"nb_points,omitempty"`
	// 近12个月内赚取的会员积分
	NbPointsEarnLast12Months int64 `json:"nb_points_earn_last_12_months,omitempty" xml:"nb_points_earn_last_12_months,omitempty"`
	// 升级所需会员积分
	NbPointsToNextTiering int64 `json:"nb_points_to_next_tiering,omitempty" xml:"nb_points_to_next_tiering,omitempty"`
	// 更新等级所获得的积分
	PointsEarnedOnTierUpdate int64 `json:"points_earned_on_tier_update,omitempty" xml:"points_earned_on_tier_update,omitempty"`
	// 当前间夜数余额
	CurrentNightsBalance int64 `json:"current_nights_balance,omitempty" xml:"current_nights_balance,omitempty"`
	// 当前入住间夜数余额
	CurrentStaysBalance int64 `json:"current_stays_balance,omitempty" xml:"current_stays_balance,omitempty"`
	// 快速跟踪间夜状态
	FastTrackedStatusNights int64 `json:"fast_tracked_status_nights,omitempty" xml:"fast_tracked_status_nights,omitempty"`
	// 升级所需会员间夜
	NbNightsToNextTiering int64 `json:"nb_nights_to_next_tiering,omitempty" xml:"nb_nights_to_next_tiering,omitempty"`
	// 更新等级所花费的间夜
	NightsSpentOnTierUpdate int64 `json:"nights_spent_on_tier_update,omitempty" xml:"nights_spent_on_tier_update,omitempty"`
}
