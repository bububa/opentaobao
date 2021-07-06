package wenyuvideo

// ShowBaseRbo 结构体
type ShowBaseRbo struct {
	// 节目名称
	ShowName string `json:"show_name,omitempty" xml:"show_name,omitempty"`
	// 节目子标题
	ShowSubtitle string `json:"show_subtitle,omitempty" xml:"show_subtitle,omitempty"`
	// 节目默认图片url
	ShowThumbUrl string `json:"show_thumb_url,omitempty" xml:"show_thumb_url,omitempty"`
	// 节目竖版图片url
	ShowVthumbUrl string `json:"show_vthumb_url,omitempty" xml:"show_vthumb_url,omitempty"`
	// 评分
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 角标
	Mark string `json:"mark,omitempty" xml:"mark,omitempty"`
	// 发行日期
	ReleaseDate string `json:"release_date,omitempty" xml:"release_date,omitempty"`
	// tag标记
	ViewTag string `json:"view_tag,omitempty" xml:"view_tag,omitempty"`
	// 影视信息在列表搜索等未知的重要提示
	Tips string `json:"tips,omitempty" xml:"tips,omitempty"`
	// 展示类型
	ShowType int64 `json:"show_type,omitempty" xml:"show_type,omitempty"`
	// 节目主分类
	ShowCategory int64 `json:"show_category,omitempty" xml:"show_category,omitempty"`
	// 是否动态更新集数
	IsDynTotal int64 `json:"is_dyn_total,omitempty" xml:"is_dyn_total,omitempty"`
	// 最新一集
	LastSequence int64 `json:"last_sequence,omitempty" xml:"last_sequence,omitempty"`
	// 正片总集数
	EpisodeTotal int64 `json:"episode_total,omitempty" xml:"episode_total,omitempty"`
	// 老媒资节目ID（整体兼容使用，不建议客户端使用）
	ProgramId int64 `json:"program_id,omitempty" xml:"program_id,omitempty"`
	// 最后一个正片集数
	EpisodeLast int64 `json:"episode_last,omitempty" xml:"episode_last,omitempty"`
	// 是否预告片
	Prevue bool `json:"prevue,omitempty" xml:"prevue,omitempty"`
}
