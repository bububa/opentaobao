package alihouse

// MarketingCouponDto 结构体
type MarketingCouponDto struct {
	// 使用楼盘列表
	ActivityProjectList []ProjectDetailInfoDto `json:"activity_project_list,omitempty" xml:"activity_project_list>project_detail_info_dto,omitempty"`
	// 券有效期
	ValidDate int64 `json:"valid_date,omitempty" xml:"valid_date,omitempty"`
	// 活动结束时间
	ActivityEndTime string `json:"activity_end_time,omitempty" xml:"activity_end_time,omitempty"`
	// 活动开始时间
	ActivityBeginTime string `json:"activity_begin_time,omitempty" xml:"activity_begin_time,omitempty"`
	// 券金额
	CouponMoney int64 `json:"coupon_money,omitempty" xml:"coupon_money,omitempty"`
	// 券状态
	ActivityStatus int64 `json:"activity_status,omitempty" xml:"activity_status,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 外部活动ID
	OuterActivityId string `json:"outer_activity_id,omitempty" xml:"outer_activity_id,omitempty"`
}
