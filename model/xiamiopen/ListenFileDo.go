package xiamiopen

import (
	"sync"
)

// ListenFileDo 结构体
type ListenFileDo struct {
	// 试听文件地址
	ListenFile string `json:"listen_file,omitempty" xml:"listen_file,omitempty"`
	// 歌曲品质，l为低品质，h为高品质，s为无损
	Quality string `json:"quality,omitempty" xml:"quality,omitempty"`
	// 超时时间
	Expire int64 `json:"expire,omitempty" xml:"expire,omitempty"`
}

var poolListenFileDo = sync.Pool{
	New: func() any {
		return new(ListenFileDo)
	},
}

// GetListenFileDo() 从对象池中获取ListenFileDo
func GetListenFileDo() *ListenFileDo {
	return poolListenFileDo.Get().(*ListenFileDo)
}

// ReleaseListenFileDo 释放ListenFileDo
func ReleaseListenFileDo(v *ListenFileDo) {
	v.ListenFile = ""
	v.Quality = ""
	v.Expire = 0
	poolListenFileDo.Put(v)
}
