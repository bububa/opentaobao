package aliexpress

import (
	"sync"
)

// OffsitePostVideoVo 结构体
type OffsitePostVideoVo struct {
	// 视频封面，与视频尺寸要一致
	CoverUrl string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// 视频地址
	VideoUrl string `json:"video_url,omitempty" xml:"video_url,omitempty"`
	// 视频宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
	// 视频高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
}

var poolOffsitePostVideoVo = sync.Pool{
	New: func() any {
		return new(OffsitePostVideoVo)
	},
}

// GetOffsitePostVideoVo() 从对象池中获取OffsitePostVideoVo
func GetOffsitePostVideoVo() *OffsitePostVideoVo {
	return poolOffsitePostVideoVo.Get().(*OffsitePostVideoVo)
}

// ReleaseOffsitePostVideoVo 释放OffsitePostVideoVo
func ReleaseOffsitePostVideoVo(v *OffsitePostVideoVo) {
	v.CoverUrl = ""
	v.VideoUrl = ""
	v.Width = 0
	v.Height = 0
	poolOffsitePostVideoVo.Put(v)
}
