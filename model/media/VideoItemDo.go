package media

import (
	"sync"
)

// VideoItemDo 结构体
type VideoItemDo struct {
	// 视频封面
	CoverUrl string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// 视频上传视频
	UploadTime string `json:"upload_time,omitempty" xml:"upload_time,omitempty"`
	// 视频标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 视频时长
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
}

var poolVideoItemDo = sync.Pool{
	New: func() any {
		return new(VideoItemDo)
	},
}

// GetVideoItemDo() 从对象池中获取VideoItemDo
func GetVideoItemDo() *VideoItemDo {
	return poolVideoItemDo.Get().(*VideoItemDo)
}

// ReleaseVideoItemDo 释放VideoItemDo
func ReleaseVideoItemDo(v *VideoItemDo) {
	v.CoverUrl = ""
	v.UploadTime = ""
	v.Title = ""
	v.Duration = 0
	poolVideoItemDo.Put(v)
}
