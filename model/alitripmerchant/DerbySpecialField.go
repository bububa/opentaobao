package alitripmerchant

// DerbySpecialField 结构体
type DerbySpecialField struct {
	// 积分倒计时
	PointsExpirationDate string `json:"points_expiration_date,omitempty" xml:"points_expiration_date,omitempty"`
	// 会员卡倒计时
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
}
