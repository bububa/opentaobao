package tbitem

import (
	"sync"
)

// MpicVideo 结构体
type MpicVideo struct {
	// 主图视频的缩略图URL
	VideoPic string `json:"video_pic,omitempty" xml:"video_pic,omitempty"`
	// 主图视频记录所关联的商品的数字id
	NumIid int64 `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 主图视频的时长，单位：秒
	VideoDuaration int64 `json:"video_duaration,omitempty" xml:"video_duaration,omitempty"`
	// 主图视频的在淘视频中的ID
	VideoId int64 `json:"video_id,omitempty" xml:"video_id,omitempty"`
	// 主图视频的状态
	VideoStatus int64 `json:"video_status,omitempty" xml:"video_status,omitempty"`
}

var poolMpicVideo = sync.Pool{
	New: func() any {
		return new(MpicVideo)
	},
}

// GetMpicVideo() 从对象池中获取MpicVideo
func GetMpicVideo() *MpicVideo {
	return poolMpicVideo.Get().(*MpicVideo)
}

// ReleaseMpicVideo 释放MpicVideo
func ReleaseMpicVideo(v *MpicVideo) {
	v.VideoPic = ""
	v.NumIid = 0
	v.VideoDuaration = 0
	v.VideoId = 0
	v.VideoStatus = 0
	poolMpicVideo.Put(v)
}
