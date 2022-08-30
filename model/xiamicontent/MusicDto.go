package xiamicontent

// MusicDto 结构体
type MusicDto struct {
	// 演唱者
	Singers []ArtistDto `json:"singers,omitempty" xml:"singers>artist_dto,omitempty"`
	// 标签
	Tags []TagLink `json:"tags,omitempty" xml:"tags>tag_link,omitempty"`
	// 音频列表
	Audios []AudioDto `json:"audios,omitempty" xml:"audios>audio_dto,omitempty"`
	// 歌词
	Lyrics []LyricDto `json:"lyrics,omitempty" xml:"lyrics>lyric_dto,omitempty"`
	// 歌曲名
	SongName string `json:"song_name,omitempty" xml:"song_name,omitempty"`
	// 封面地址
	CoverUrl string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// 波形图地址
	WaveformUrl string `json:"waveform_url,omitempty" xml:"waveform_url,omitempty"`
	// 副标题
	SongSubName string `json:"song_sub_name,omitempty" xml:"song_sub_name,omitempty"`
	// 推荐场景id
	RecommendSceneId string `json:"recommend_scene_id,omitempty" xml:"recommend_scene_id,omitempty"`
	// 专辑
	Album *AlbumDto `json:"album,omitempty" xml:"album,omitempty"`
	// 展示状态
	ShowStatus int64 `json:"show_status,omitempty" xml:"show_status,omitempty"`
	// 时长
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 1-版权上架, 0-版权下架
	CopyrightStatus int64 `json:"copyright_status,omitempty" xml:"copyright_status,omitempty"`
	// 歌曲id
	SongId int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
	// 音乐类型1完整版2剪辑版
	MusicType int64 `json:"music_type,omitempty" xml:"music_type,omitempty"`
	// 发布状态
	PublishStatus int64 `json:"publish_status,omitempty" xml:"publish_status,omitempty"`
	// 是否删除
	DeletedStatus int64 `json:"deleted_status,omitempty" xml:"deleted_status,omitempty"`
}
