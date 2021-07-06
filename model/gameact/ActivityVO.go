package gameact

// ActivityVO 结构体
type ActivityVO struct {
	// 奖项列表
	Awards []AwardVO `json:"awards,omitempty" xml:"awards>award_vo,omitempty"`
	// 活动连接
	ActivityUrl string `json:"activity_url,omitempty" xml:"activity_url,omitempty"`
	// 活动名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 活动描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 运营和cp约定的唯一事件标示
	EventKey string `json:"event_key,omitempty" xml:"event_key,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 1970年到现在的毫秒数
	StartTime int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 1970年距离现在的毫秒数
	EndTime int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 积分/金牌消耗
	ConsumeAmount int64 `json:"consume_amount,omitempty" xml:"consume_amount,omitempty"`
	// 抽奖类型
	LuckyType int64 `json:"lucky_type,omitempty" xml:"lucky_type,omitempty"`
	// 抽奖渠道
	LuckyChannel int64 `json:"lucky_channel,omitempty" xml:"lucky_channel,omitempty"`
	// 抽奖次数（免费）
	AccessAmount int64 `json:"access_amount,omitempty" xml:"access_amount,omitempty"`
}
