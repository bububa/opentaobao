package xiamiopen

import (
	"sync"
)

// AlbumDo 结构体
type AlbumDo struct {
	// 专辑名称
	AlbumName string `json:"album_name,omitempty" xml:"album_name,omitempty"`
}

var poolAlbumDo = sync.Pool{
	New: func() any {
		return new(AlbumDo)
	},
}

// GetAlbumDo() 从对象池中获取AlbumDo
func GetAlbumDo() *AlbumDo {
	return poolAlbumDo.Get().(*AlbumDo)
}

// ReleaseAlbumDo 释放AlbumDo
func ReleaseAlbumDo(v *AlbumDo) {
	v.AlbumName = ""
	poolAlbumDo.Put(v)
}
