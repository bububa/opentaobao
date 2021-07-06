package caipiao

// WangcaiMarketingDetail 结构体
type WangcaiMarketingDetail struct {
	// 参与活动的商品ID
	Items []int64 `json:"items,omitempty" xml:"items>int64,omitempty"`
	// 外部系统主键
	BizId string `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 活动ID,一个活动可以关联多个送彩票设置
	ActivityId string `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动开始时间
	BeginTime string `json:"begin_time,omitempty" xml:"begin_time,omitempty"`
	// 活动结束时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 活动类型,0全店/1指定商品
	ActivityType int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 活动的最小金额门槛,以分为单位
	MinAmount int64 `json:"min_amount,omitempty" xml:"min_amount,omitempty"`
	// 赠送的彩种,1双色球/8大乐透
	LotteryTypeId int64 `json:"lottery_type_id,omitempty" xml:"lottery_type_id,omitempty"`
	// 赠送的彩票注数
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}
