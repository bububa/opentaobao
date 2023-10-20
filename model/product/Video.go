package product

import (
	"sync"
)

// Video 结构体
type Video struct {
	// video的url连接地址。淘秀里视频记录里面存储的url地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 视频关联记录创建时间（格式：yyyy-MM-dd HH:mm:ss）
	Created string `json:"created,omitempty" xml:"created,omitempty"`
	// 视频关联记录修改时间（格式：yyyy-MM-dd HH:mm:ss）
	Modified string `json:"modified,omitempty" xml:"modified,omitempty"`
}

var poolVideo = sync.Pool{
	New: func() any {
		return new(Video)
	},
}

// GetVideo() 从对象池中获取Video
func GetVideo() *Video {
	return poolVideo.Get().(*Video)
}

// ReleaseVideo 释放Video
func ReleaseVideo(v *Video) {
	v.Url = ""
	v.Created = ""
	v.Modified = ""
	poolVideo.Put(v)
}
