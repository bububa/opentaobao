package xiami

// Songs 结构体
type Songs struct {
	// songName
	SongName string `json:"song_name,omitempty" xml:"song_name,omitempty"`
	// albumName
	AlbumName string `json:"album_name,omitempty" xml:"album_name,omitempty"`
	// artistName
	ArtistName string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
	// 专辑id
	AlbumId int64 `json:"album_id,omitempty" xml:"album_id,omitempty"`
	// 歌曲id
	SongId int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
	// 0：不可试听，1：可试听
	IsPlay int64 `json:"is_play,omitempty" xml:"is_play,omitempty"`
	// 下载标识，1表示不允许下载
	Flag int64 `json:"flag,omitempty" xml:"flag,omitempty"`
	// 歌手
	Singers string `json:"singers,omitempty" xml:"singers,omitempty"`
	// 歌曲时长
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
	// 艺人id
	ArtistId int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
	// 专辑封面
	AlbumLogo string `json:"album_logo,omitempty" xml:"album_logo,omitempty"`
	// mv_id
	MvId string `json:"mv_id,omitempty" xml:"mv_id,omitempty"`
}
