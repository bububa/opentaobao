package tmallgenie

// RawVideoAlbum 结构体
type RawVideoAlbum struct {
	// 演员名称
	ActorName []string `json:"actor_name,omitempty" xml:"actor_name>string,omitempty"`
	// 别名
	Alias []string `json:"alias,omitempty" xml:"alias>string,omitempty"`
	// 导演名称
	DirectorName []string `json:"director_name,omitempty" xml:"director_name>string,omitempty"`
	// 上传者名称
	ProducerName []string `json:"producer_name,omitempty" xml:"producer_name>string,omitempty"`
	// 标签ID，具体取值范围参见文档说明
	TagIds []int64 `json:"tag_ids,omitempty" xml:"tag_ids>int64,omitempty"`
	// 上传者名称
	UploaderName []string `json:"uploader_name,omitempty" xml:"uploader_name>string,omitempty"`
	// 专辑封面(竖图296 * 440，根据搜索规则，提供竖图用户搜索时可优先搜索到此内容)
	VCoverUrl string `json:"v_cover_url,omitempty" xml:"v_cover_url,omitempty"`
	// 区域
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 评分
	ContentScore string `json:"content_score,omitempty" xml:"content_score,omitempty"`
	// 专辑封面(横图，图片尺寸是295 * 167)
	CoverUrl string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// 描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 扩展字段
	ExtendInfo string `json:"extend_info,omitempty" xml:"extend_info,omitempty"`
	// 三方视频辑ID
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 语言
	Language string `json:"language,omitempty" xml:"language,omitempty"`
	// 操作类型 ADD/UPDATE/DELETE
	Operation string `json:"operation,omitempty" xml:"operation,omitempty"`
	// 视频来源类型，PGC/UGC/OGC
	OupgcType string `json:"oupgc_type,omitempty" xml:"oupgc_type,omitempty"`
	// 子标题
	SubTitle string `json:"sub_title,omitempty" xml:"sub_title,omitempty"`
	// 标题（若需显示集数，请填写：第X集：视频名称；如第一集：天猫精灵）
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 试看数(表示前N个视频支持试看)
	AuditionNum int64 `json:"audition_num,omitempty" xml:"audition_num,omitempty"`
	// 付费类型 0-免费、1-VIP免费、2-整本专辑售卖
	ChargeType int64 `json:"charge_type,omitempty" xml:"charge_type,omitempty"`
	// 收藏量
	CollectCount int64 `json:"collect_count,omitempty" xml:"collect_count,omitempty"`
	// 评论量
	CommentCount int64 `json:"comment_count,omitempty" xml:"comment_count,omitempty"`
	// 类目ID，参见相关文档说明
	CommonCateId int64 `json:"common_cate_id,omitempty" xml:"common_cate_id,omitempty"`
	// 成本价(单位:分)
	CostPrice int64 `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
	// 转发量
	ForwardCount int64 `json:"forward_count,omitempty" xml:"forward_count,omitempty"`
	// 是否独播 0-非独播 1-独播
	IsExclusive int64 `json:"is_exclusive,omitempty" xml:"is_exclusive,omitempty"`
	// 是否完结，0-否，1-是，默认0
	IsFinished int64 `json:"is_finished,omitempty" xml:"is_finished,omitempty"`
	// 点赞数
	LikeCount int64 `json:"like_count,omitempty" xml:"like_count,omitempty"`
	// 播放量
	PlayCount int64 `json:"play_count,omitempty" xml:"play_count,omitempty"`
	// 播放顺序0-顺序，1-倒序
	PlayOrder int64 `json:"play_order,omitempty" xml:"play_order,omitempty"`
	// 发布时间,unix时间戳,单位:毫秒
	ReleaseTime int64 `json:"release_time,omitempty" xml:"release_time,omitempty"`
	// 建议最高零售价(单位:分)
	SuggestMaxPrice int64 `json:"suggest_max_price,omitempty" xml:"suggest_max_price,omitempty"`
	// 建议最低零售价(单位:分)
	SuggestMinPrice int64 `json:"suggest_min_price,omitempty" xml:"suggest_min_price,omitempty"`
	// 包含视频总集数
	TotalEpisode int64 `json:"total_episode,omitempty" xml:"total_episode,omitempty"`
	// 失效时间,unix时间戳,单位:毫秒,为空则永不失效
	ValidEndTime int64 `json:"valid_end_time,omitempty" xml:"valid_end_time,omitempty"`
	// 生效时间,unix时间戳,单位:毫秒,为空则立即生效
	ValidStartTime int64 `json:"valid_start_time,omitempty" xml:"valid_start_time,omitempty"`
	// 更新到第几集，如果是未完结状态此字段必填
	UpdateIndex int64 `json:"update_index,omitempty" xml:"update_index,omitempty"`
}
