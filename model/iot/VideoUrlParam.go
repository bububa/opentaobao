package iot

import (
	"sync"
)

// VideoUrlParam 结构体
type VideoUrlParam struct {
	// 默认播放链接
	DefaultUrl string `json:"default_url,omitempty" xml:"default_url,omitempty"`
	// 高清
	High string `json:"high,omitempty" xml:"high,omitempty"`
	// 标清
	Standard string `json:"standard,omitempty" xml:"standard,omitempty"`
	// 超清
	Ultra string `json:"ultra,omitempty" xml:"ultra,omitempty"`
	// 视频封面图
	Cover *ImageUrlParam `json:"cover,omitempty" xml:"cover,omitempty"`
	// 视频高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 视频宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}

var poolVideoUrlParam = sync.Pool{
	New: func() any {
		return new(VideoUrlParam)
	},
}

// GetVideoUrlParam() 从对象池中获取VideoUrlParam
func GetVideoUrlParam() *VideoUrlParam {
	return poolVideoUrlParam.Get().(*VideoUrlParam)
}

// ReleaseVideoUrlParam 释放VideoUrlParam
func ReleaseVideoUrlParam(v *VideoUrlParam) {
	v.DefaultUrl = ""
	v.High = ""
	v.Standard = ""
	v.Ultra = ""
	v.Cover = nil
	v.Height = 0
	v.Width = 0
	poolVideoUrlParam.Put(v)
}
