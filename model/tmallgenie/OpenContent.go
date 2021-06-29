package tmallgenie

// OpenContent 
type OpenContent struct {
    // 内容在原平台所属专辑或分类的id
    AlbumRawId   int64 `json:"album_raw_id,omitempty" xml:"album_raw_id,omitempty"`
    // 内容标签
    Tags   []string `json:"tags,omitempty" xml:"tags>string,omitempty"`
    // 内容在原平台的id，当只推送专辑信息时，此值可填0
    RawId   int64 `json:"raw_id,omitempty" xml:"raw_id,omitempty"`
    // 封面图片信息
    ImageUrl   *ImageUrl `json:"image_url,omitempty" xml:"image_url,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 内容所属专辑或分类名称/标题,这里如果专辑标题是空则不同步专辑信息
    AlbumTitle   string `json:"album_title,omitempty" xml:"album_title,omitempty"`
    // 其他扩展字段，不同类型内容有不同要求，具体请参加详细说明文档extend_info字段部分
    ExtendInfo   string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
    // 操作方式，支持新增和删除操作(ADD/DELETE)
    Operation   string `json:"operation,omitempty" xml:"operation,omitempty"`
    // 作者信息
    Author   *Author `json:"author,omitempty" xml:"author,omitempty"`
    // 内容标题或者名称，最长150个字符
    Title   string `json:"title,omitempty" xml:"title,omitempty"`
    // 播放时长(单位：秒)
    Duration   int64 `json:"duration,omitempty" xml:"duration,omitempty"`
    // 内容在原平台播放次数
    PlayCount   int64 `json:"play_count,omitempty" xml:"play_count,omitempty"`
    // 内容描述信息(限长500)
    Description   string `json:"description,omitempty" xml:"description,omitempty"`
    // 内容发布时间，1970年1月1日至当前的秒数(unix时间戳)
    ReleaseTime   int64 `json:"release_time,omitempty" xml:"release_time,omitempty"`
    // 播放链接
    PlayUrls   []PlayUrl `json:"play_urls,omitempty" xml:"play_urls>play_url,omitempty"`
    // 内容所属专辑或分类描述信息
    AlbumDescription   string `json:"album_description,omitempty" xml:"album_description,omitempty"`
    // 在有所属专辑/分类的场景下，标识此音频在专辑/分类下的顺序值，从1开始
    SortNum   int64 `json:"sort_num,omitempty" xml:"sort_num,omitempty"`
    // 描述类型,如果描述是TTS类型会做TTS处理
    DescriptionType   string `json:"description_type,omitempty" xml:"description_type,omitempty"`
    // 三方热度分1~5
    HotScore   int64 `json:"hot_score,omitempty" xml:"hot_score,omitempty"`
    // 0-免费、1-单篇售卖，2-专辑售卖，3-商业化一期类型
    ChargeType   int64 `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
    // 播放顺序 0 正序 1 倒序
    PlayOrder   int64 `json:"play_order,omitempty" xml:"play_order,omitempty"`
    // 产品描述信息
    ProductDesc   string `json:"product_desc,omitempty" xml:"product_desc,omitempty"`
    // 标签id见智能应用平台
    TagIds   []int64 `json:"tag_ids,omitempty" xml:"tag_ids>int64,omitempty"`
    // 是否vip付费
    VipFree   bool `json:"vip_free,omitempty" xml:"vip_free,omitempty"`
    // 成本价(单位分)，没有明确区分就成本最高最低传一样的值
    CostPrice   int64 `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
    // 最小建议零售价(单位分)，没有明确区分就成本最高最低传一样的值
    SuggestMinPrice   int64 `json:"suggest_min_price,omitempty" xml:"suggest_min_price,omitempty"`
    // 最大建议零售价(单位分)，没有明确区分就成本最高最低传一样的值
    SuggestMaxPrice   int64 `json:"suggest_max_price,omitempty" xml:"suggest_max_price,omitempty"`
    // 专辑是否支持试听,默认为false
    IsAlbumAudition   bool `json:"is_album_audition,omitempty" xml:"is_album_audition,omitempty"`
    // 内容分集是否支持试听,默认为false
    IsAudition   bool `json:"is_audition,omitempty" xml:"is_audition,omitempty"`
}
