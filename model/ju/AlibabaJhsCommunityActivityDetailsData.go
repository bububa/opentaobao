package ju

// AlibabajhscommunityactivitydetailsData 结构体
type AlibabajhscommunityactivitydetailsData struct {
	// 开幕时间
	OpeningTime string `json:"opening_time,omitempty" xml:"opening_time,omitempty"`
	// 淘口令
	TaoToken string `json:"tao_token,omitempty" xml:"tao_token,omitempty"`
	// 是否已预约
	Reservation string `json:"reservation,omitempty" xml:"reservation,omitempty"`
	// 活动id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 活动标题
	ActivityTitle string `json:"activity_title,omitempty" xml:"activity_title,omitempty"`
	// 活动内容
	ActivityContent string `json:"activity_content,omitempty" xml:"activity_content,omitempty"`
	// 活动回流地址
	ActivityBackUrl string `json:"activity_back_url,omitempty" xml:"activity_back_url,omitempty"`
	// 渠道外传播标题
	SpreadTitle string `json:"spread_title,omitempty" xml:"spread_title,omitempty"`
	// 对外传播图片
	SpreadPicUrl string `json:"spread_pic_url,omitempty" xml:"spread_pic_url,omitempty"`
	// 活动状态
	ActivityStatus string `json:"activity_status,omitempty" xml:"activity_status,omitempty"`
	// 活动类型
	ActivityType int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
}
