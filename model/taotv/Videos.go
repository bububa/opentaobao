package taotv

import (
	"sync"
)

// Videos 结构体
type Videos struct {
	// 视频标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 视频ID
	VideoId string `json:"video_id,omitempty" xml:"video_id,omitempty"`
	// 视频图片
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 视频来源
	From string `json:"from,omitempty" xml:"from,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolVideos = sync.Pool{
	New: func() any {
		return new(Videos)
	},
}

// GetVideos() 从对象池中获取Videos
func GetVideos() *Videos {
	return poolVideos.Get().(*Videos)
}

// ReleaseVideos 释放Videos
func ReleaseVideos(v *Videos) {
	v.Title = ""
	v.VideoId = ""
	v.PicUrl = ""
	v.From = ""
	v.Id = 0
	poolVideos.Put(v)
}
