package tanx

// NativeTemplateCreativeDto 
type NativeTemplateCreativeDto struct {

    // 多选一的属性集合，这些属性至少有一个不为空，1:标题;2:广告语;3:图片;4:价格;5:折扣价;6:销量;7:click_url;8:landing_type;9描述;10打开方式;11下载方式;12deepLink;13下载
    MutlichoiceFields   []Number `json:"mutlichoice_fields,omitempty"`

    // 必须属性集合，1:标题;2:广告语;3:图片;4:价格;5:折扣价;6:销量;7:click_url;8:landing_type;9描述;10打开方式;11下载方式;12deepLink;13下载
    RequiredFields   []Number `json:"required_fields,omitempty"`

    // 副标题或者广告语最大长度，如果超长有可能会被截断
    AdWordsMaxSafeLength   int64 `json:"ad_words_max_safe_length,omitempty"`

    // 推荐属性集合，1:标题;2:广告语;3:图片;4:价格;5:折扣价;6:销量;7:click_url;8:landing_type;9描述;10打开方式;11下载方式;12deepLink;13下载
    RecommendedFields   []Number `json:"recommended_fields,omitempty"`

    // 图片尺寸宽x高，比如320x50
    ImageSize   string `json:"image_size,omitempty"`

    // 标题最大长度，如果超长有可能会被截断
    TitleMaxSafeLength   int64 `json:"title_max_safe_length,omitempty"`

    // 已经废弃，事件属性集合 1:下载 2:地图 3:电话
    ActionFields   []Number `json:"action_fields,omitempty"`

}
