package xiami

// RankSong 结构体
type RankSong struct {
	// 歌曲ID
	SongId int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
	// 歌曲名称
	SongName string `json:"song_name,omitempty" xml:"song_name,omitempty"`
	// 专辑ID
	AlbumId int64 `json:"album_id,omitempty" xml:"album_id,omitempty"`
	// 专辑名
	AlbumName string `json:"album_name,omitempty" xml:"album_name,omitempty"`
	// 艺人ID
	ArtistId int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
	// 艺人名
	ArtistName string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
	// 文本歌词
	LyricText string `json:"lyric_text,omitempty" xml:"lyric_text,omitempty"`
	// 推荐值
	Recommends int64 `json:"recommends,omitempty" xml:"recommends,omitempty"`
	// 歌曲时长（S）
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 专辑LOGO
	Logo string `json:"logo,omitempty" xml:"logo,omitempty"`
	// 艺人LOGO
	ArtistLogo string `json:"artist_logo,omitempty" xml:"artist_logo,omitempty"`
	// 演唱者
	Singers string `json:"singers,omitempty" xml:"singers,omitempty"`
	// 播放次数
	PlayCounts int64 `json:"play_counts,omitempty" xml:"play_counts,omitempty"`
	// 音频地址
	ListenFile string `json:"listen_file,omitempty" xml:"listen_file,omitempty"`
	// 播放时长 （同length)
	PlaySeconds int64 `json:"play_seconds,omitempty" xml:"play_seconds,omitempty"`
	// 专辑名称(同album_name)
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 歌曲名称（同song_name）
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 专辑LOGO(同logo)
	AlbumLogo string `json:"album_logo,omitempty" xml:"album_logo,omitempty"`
	// 动态歌词（同lyric）
	LyricFile string `json:"lyric_file,omitempty" xml:"lyric_file,omitempty"`
	// 是否音乐人demo(0,不是,1,是)
	Demo int64 `json:"demo,omitempty" xml:"demo,omitempty"`
	// 播放权限(1,可以播放,0,不可播放)
	PlayAuthority int64 `json:"play_authority,omitempty" xml:"play_authority,omitempty"`
	// 是否已收藏
	Favourite bool `json:"favourite,omitempty" xml:"favourite,omitempty"`
	// 试听占比, 某些榜单会有
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}
