package youkuott

// YoukutvoperatormediapagequeryData 结构体
type YoukutvoperatormediapagequeryData struct {
	// 分类列表
	GenreList []string `json:"genre_list,omitempty" xml:"genre_list>string,omitempty"`
	// 导演
	DirectorList []string `json:"director_list,omitempty" xml:"director_list>string,omitempty"`
	// 演员
	PerformerList []string `json:"performer_list,omitempty" xml:"performer_list>string,omitempty"`
	// 节目国家地区列表
	AreaList []string `json:"area_list,omitempty" xml:"area_list>string,omitempty"`
	// 优酷标签
	YoukuTags string `json:"youku_tags,omitempty" xml:"youku_tags,omitempty"`
	// 分数
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 子标题
	ShowSubtitle string `json:"show_subtitle,omitempty" xml:"show_subtitle,omitempty"`
	// 节目banner
	ShowBannerUrl string `json:"show_banner_url,omitempty" xml:"show_banner_url,omitempty"`
	// 分类名称
	ShowCategoryName string `json:"show_category_name,omitempty" xml:"show_category_name,omitempty"`
	// 纵向海报
	ShowVthumbUrl string `json:"show_vthumb_url,omitempty" xml:"show_vthumb_url,omitempty"`
	// 节目描述
	ShowDesc string `json:"show_desc,omitempty" xml:"show_desc,omitempty"`
	// 节目名称
	ShowName string `json:"show_name,omitempty" xml:"show_name,omitempty"`
	// 节目id
	ShowId string `json:"show_id,omitempty" xml:"show_id,omitempty"`
	// 横版海报
	ShowThumbUrl string `json:"show_thumb_url,omitempty" xml:"show_thumb_url,omitempty"`
	// 更新集数
	EpisodeLast int64 `json:"episode_last,omitempty" xml:"episode_last,omitempty"`
	// 是否收费，0=免费，1=收费
	Paid int64 `json:"paid,omitempty" xml:"paid,omitempty"`
	// 发行日期
	ReleaseDate int64 `json:"release_date,omitempty" xml:"release_date,omitempty"`
	// 总集数
	EpisodeTotal int64 `json:"episode_total,omitempty" xml:"episode_total,omitempty"`
	// 码流二进制组合
	PlaySet int64 `json:"play_set,omitempty" xml:"play_set,omitempty"`
	// 时长
	Seconds int64 `json:"seconds,omitempty" xml:"seconds,omitempty"`
	// 媒资最后修改时间
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 上线下状态
	VmacState int64 `json:"vmac_state,omitempty" xml:"vmac_state,omitempty"`
}
