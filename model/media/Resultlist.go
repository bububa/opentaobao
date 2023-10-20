package media

import (
	"sync"
)

// Resultlist 结构体
type Resultlist struct {
	// 视频链接，需要unicode转换
	VideoUrl string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	// 封面图url
	CoverUrl string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// 视频标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 互动id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 视频时长 单位秒
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 视频状态，1有效；0删除
	VideoStatus int64 `json:"video_status,omitempty" xml:"video_status,omitempty"`
}

var poolResultlist = sync.Pool{
	New: func() any {
		return new(Resultlist)
	},
}

// GetResultlist() 从对象池中获取Resultlist
func GetResultlist() *Resultlist {
	return poolResultlist.Get().(*Resultlist)
}

// ReleaseResultlist 释放Resultlist
func ReleaseResultlist(v *Resultlist) {
	v.VideoUrl = ""
	v.CoverUrl = ""
	v.Title = ""
	v.Id = 0
	v.Duration = 0
	v.VideoStatus = 0
	poolResultlist.Put(v)
}
