package tanx

// NativeTemplateCreativeDto 结构体
type NativeTemplateCreativeDto struct {
	// 多选一的属性集合，这些属性至少有一个不为空，1:标题;2:广告语;3:图片;4:价格;5:折扣价;6:销量;7:click_url;8:landing_type;9描述;10打开方式;11下载方式;12deepLink;13下载
	MutlichoiceFields []int64 `json:"mutlichoice_fields,omitempty" xml:"mutlichoice_fields>int64,omitempty"`
	// 必须属性集合，1:标题;2:广告语;3:图片;4:价格;5:折扣价;6:销量;7:click_url;8:landing_type;9描述;10打开方式;11下载方式;12deepLink;13下载
	RequiredFields []int64 `json:"required_fields,omitempty" xml:"required_fields>int64,omitempty"`
	// 推荐属性集合，1:标题;2:广告语;3:图片;4:价格;5:折扣价;6:销量;7:click_url;8:landing_type;9描述;10打开方式;11下载方式;12deepLink;13下载
	RecommendedFields []int64 `json:"recommended_fields,omitempty" xml:"recommended_fields>int64,omitempty"`
	// 已经废弃，事件属性集合 1:下载 2:地图 3:电话
	ActionFields []int64 `json:"action_fields,omitempty" xml:"action_fields>int64,omitempty"`
	// 图片尺寸宽x高，比如320x50
	ImageSize string `json:"image_size,omitempty" xml:"image_size,omitempty"`
	// 副标题或者广告语最大长度，如果超长有可能会被截断
	AdWordsMaxSafeLength int64 `json:"ad_words_max_safe_length,omitempty" xml:"ad_words_max_safe_length,omitempty"`
	// 标题最大长度，如果超长有可能会被截断
	TitleMaxSafeLength int64 `json:"title_max_safe_length,omitempty" xml:"title_max_safe_length,omitempty"`
}
