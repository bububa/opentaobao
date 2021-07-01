package taotv

// V5BaseItemRbo 结构体
type V5BaseItemRbo struct {
	// 坑位ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 坑位标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 坑位子标题
	Subtitle string `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	// 中间图片
	CenterPic string `json:"center_pic,omitempty" xml:"center_pic,omitempty"`
	// 背景图片
	BgPic string `json:"bg_pic,omitempty" xml:"bg_pic,omitempty"`
	// 背景动图（高配）
	BgPicGif string `json:"bg_pic_gif,omitempty" xml:"bg_pic_gif,omitempty"`
	// 跳转行为
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 扩展字段
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 未解析的扩展字段
	ExtraStr string `json:"extra_str,omitempty" xml:"extra_str,omitempty"`
	// 坑位类型
	ItemType string `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 评分
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 推荐主题ID
	RecommendRuleId int64 `json:"recommend_rule_id,omitempty" xml:"recommend_rule_id,omitempty"`
	// 推荐理由
	RecommendReason string `json:"recommend_reason,omitempty" xml:"recommend_reason,omitempty"`
}
