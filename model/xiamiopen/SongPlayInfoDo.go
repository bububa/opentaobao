package xiamiopen

import (
	"sync"
)

// SongPlayInfoDo 结构体
type SongPlayInfoDo struct {
	// 试听文件列表
	ListenFileList []ListenFileDo `json:"listen_file_list,omitempty" xml:"listen_file_list>listen_file_do,omitempty"`
	// 歌曲id
	SongId int64 `json:"song_id,omitempty" xml:"song_id,omitempty"`
}

var poolSongPlayInfoDo = sync.Pool{
	New: func() any {
		return new(SongPlayInfoDo)
	},
}

// GetSongPlayInfoDo() 从对象池中获取SongPlayInfoDo
func GetSongPlayInfoDo() *SongPlayInfoDo {
	return poolSongPlayInfoDo.Get().(*SongPlayInfoDo)
}

// ReleaseSongPlayInfoDo 释放SongPlayInfoDo
func ReleaseSongPlayInfoDo(v *SongPlayInfoDo) {
	v.ListenFileList = v.ListenFileList[:0]
	v.SongId = 0
	poolSongPlayInfoDo.Put(v)
}
