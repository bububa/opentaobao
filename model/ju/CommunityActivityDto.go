package ju

// CommunityActivityDto 结构体
type CommunityActivityDto struct {
	// 活动标题
	ActivityTitle string `json:"activity_title,omitempty" xml:"activity_title,omitempty"`
	// 活动内容
	ActivityContent string `json:"activity_content,omitempty" xml:"activity_content,omitempty"`
	// 回流地址
	ActivityBackUrl string `json:"activity_back_url,omitempty" xml:"activity_back_url,omitempty"`
	// 标题
	SpreadTitle string `json:"spread_title,omitempty" xml:"spread_title,omitempty"`
	// 传播图片信息
	SpreadPicUrl string `json:"spread_pic_url,omitempty" xml:"spread_pic_url,omitempty"`
	// 操作时间
	OpeningTime int64 `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 预约时间
	ReservationTime int64 `json:"reservation_time,omitempty" xml:"reservation_time,omitempty"`
	// 活动id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
