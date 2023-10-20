package taotv

import (
	"sync"
)

// Channels 结构体
type Channels struct {
	// 频道所有的视频列表
	VideoList []Videolist `json:"video_list,omitempty" xml:"video_list>videolist,omitempty"`
	// 频道描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 频道图标
	Pic string `json:"pic,omitempty" xml:"pic,omitempty"`
	// 频道名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 频道ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 频道固定的编号
	SerialNumber int64 `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
	// 频道当前播放视频
	CurrentVideo *CarouselCurrentPlayRbo `json:"current_video,omitempty" xml:"current_video,omitempty"`
	// 牌照方
	Bcp int64 `json:"bcp,omitempty" xml:"bcp,omitempty"`
}

var poolChannels = sync.Pool{
	New: func() any {
		return new(Channels)
	},
}

// GetChannels() 从对象池中获取Channels
func GetChannels() *Channels {
	return poolChannels.Get().(*Channels)
}

// ReleaseChannels 释放Channels
func ReleaseChannels(v *Channels) {
	v.VideoList = v.VideoList[:0]
	v.Description = ""
	v.Pic = ""
	v.Name = ""
	v.Id = 0
	v.SerialNumber = 0
	v.CurrentVideo = nil
	v.Bcp = 0
	poolChannels.Put(v)
}
