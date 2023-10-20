package taotv

import (
	"sync"
)

// CarouselCurrentPlayRbo 结构体
type CarouselCurrentPlayRbo struct {
	// 当前轮播视频信息
	Video *CarouselPlaylistVideoRbo `json:"video,omitempty" xml:"video,omitempty"`
	// 当前视频正在播放的时间点（单位秒）
	Point int64 `json:"point,omitempty" xml:"point,omitempty"`
}

var poolCarouselCurrentPlayRbo = sync.Pool{
	New: func() any {
		return new(CarouselCurrentPlayRbo)
	},
}

// GetCarouselCurrentPlayRbo() 从对象池中获取CarouselCurrentPlayRbo
func GetCarouselCurrentPlayRbo() *CarouselCurrentPlayRbo {
	return poolCarouselCurrentPlayRbo.Get().(*CarouselCurrentPlayRbo)
}

// ReleaseCarouselCurrentPlayRbo 释放CarouselCurrentPlayRbo
func ReleaseCarouselCurrentPlayRbo(v *CarouselCurrentPlayRbo) {
	v.Video = nil
	v.Point = 0
	poolCarouselCurrentPlayRbo.Put(v)
}
