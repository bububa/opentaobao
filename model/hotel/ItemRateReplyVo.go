package hotel

// ItemRateReplyVo 结构体
type ItemRateReplyVo struct {
	// 内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 评论时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// intervalDay
	IntervalDay int64 `json:"interval_day,omitempty" xml:"interval_day,omitempty"`
	// 图片信息,图片URL的list
	MediaInfo string `json:"media_info,omitempty" xml:"media_info,omitempty"`
	// 回复的是那一条，如果是回复主评为0，否则为追评id，组成树形结构
	ParentId int64 `json:"parent_id,omitempty" xml:"parent_id,omitempty"`
	// 被回复人的冗余信息
	ParentInfo *ParentInfo `json:"parent_info,omitempty" xml:"parent_info,omitempty"`
	// replyId
	ReplyId int64 `json:"reply_id,omitempty" xml:"reply_id,omitempty"`
	// 恢复类型
	ReplyType int64 `json:"reply_type,omitempty" xml:"reply_type,omitempty"`
	// 脱敏后的userId
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 脱敏后的用户名字
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
}
