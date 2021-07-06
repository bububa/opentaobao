package westcrm

// AlibabaWestcrmActivityInfoGetData 结构体
type AlibabaWestcrmActivityInfoGetData struct {
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 活动内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 规则
	RulesContent string `json:"rules_content,omitempty" xml:"rules_content,omitempty"`
	// 图片路径
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 活动详情路径
	DisplayUrl string `json:"display_url,omitempty" xml:"display_url,omitempty"`
	// 内容类型
	ContentType int64 `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// 活动类型
	ActivityType int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
