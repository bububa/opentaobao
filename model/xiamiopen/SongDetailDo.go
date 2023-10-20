package xiamiopen

import (
	"sync"
)

// SongDetailDo 结构体
type SongDetailDo struct {
	// 演唱者
	Singers []ArtistDo `json:"singers,omitempty" xml:"singers>artist_do,omitempty"`
	// mvId
	MvId string `json:"mv_id,omitempty" xml:"mv_id,omitempty"`
	// 专辑logo
	AlbumLogo string `json:"album_logo,omitempty" xml:"album_logo,omitempty"`
	// 编曲
	Arrangement string `json:"arrangement,omitempty" xml:"arrangement,omitempty"`
	// 作词
	Composer string `json:"composer,omitempty" xml:"composer,omitempty"`
	// 歌曲名
	SongName string `json:"song_name,omitempty" xml:"song_name,omitempty"`
	// 作曲
	Songwriters string `json:"songwriters,omitempty" xml:"songwriters,omitempty"`
	// 歌词文件地址
	LyricFile string `json:"lyric_file,omitempty" xml:"lyric_file,omitempty"`
	// 版权信息
	PurviewInfo string `json:"purview_info,omitempty" xml:"purview_info,omitempty"`
	// 专辑信息
	Album *AlbumDo `json:"album,omitempty" xml:"album,omitempty"`
	// 专辑id
	AlbumId int64 `json:"album_id,omitempty" xml:"album_id,omitempty"`
	// artistDO
	Artist *ArtistDo `json:"artist,omitempty" xml:"artist,omitempty"`
	// 艺人id
	ArtistId int64 `json:"artist_id,omitempty" xml:"artist_id,omitempty"`
	// cd序号
	CdSerial int64 `json:"cd_serial,omitempty" xml:"cd_serial,omitempty"`
	// 歌曲id
	SongId int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
	// 序号
	Track int64 `json:"track,omitempty" xml:"track,omitempty"`
	// 歌词类型
	LyricType int64 `json:"lyric_type,omitempty" xml:"lyric_type,omitempty"`
	// 歌曲时长
	Length int64 `json:"length,omitempty" xml:"length,omitempty"`
}

var poolSongDetailDo = sync.Pool{
	New: func() any {
		return new(SongDetailDo)
	},
}

// GetSongDetailDo() 从对象池中获取SongDetailDo
func GetSongDetailDo() *SongDetailDo {
	return poolSongDetailDo.Get().(*SongDetailDo)
}

// ReleaseSongDetailDo 释放SongDetailDo
func ReleaseSongDetailDo(v *SongDetailDo) {
	v.Singers = v.Singers[:0]
	v.MvId = ""
	v.AlbumLogo = ""
	v.Arrangement = ""
	v.Composer = ""
	v.SongName = ""
	v.Songwriters = ""
	v.LyricFile = ""
	v.PurviewInfo = ""
	v.Album = nil
	v.AlbumId = 0
	v.Artist = nil
	v.ArtistId = 0
	v.CdSerial = 0
	v.SongId = 0
	v.Track = 0
	v.LyricType = 0
	v.Length = 0
	poolSongDetailDo.Put(v)
}
