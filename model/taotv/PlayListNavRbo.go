package taotv

import (
	"sync"
)

// PlayListNavRbo 结构体
type PlayListNavRbo struct {
	// 当前播单视频列表
	Videos []Videos `json:"videos,omitempty" xml:"videos>videos,omitempty"`
	// 播单列表
	PlayList []Playlist `json:"play_list,omitempty" xml:"play_list>playlist,omitempty"`
	// 当前播单id
	CurPlayListId string `json:"cur_play_list_id,omitempty" xml:"cur_play_list_id,omitempty"`
}

var poolPlayListNavRbo = sync.Pool{
	New: func() any {
		return new(PlayListNavRbo)
	},
}

// GetPlayListNavRbo() 从对象池中获取PlayListNavRbo
func GetPlayListNavRbo() *PlayListNavRbo {
	return poolPlayListNavRbo.Get().(*PlayListNavRbo)
}

// ReleasePlayListNavRbo 释放PlayListNavRbo
func ReleasePlayListNavRbo(v *PlayListNavRbo) {
	v.Videos = v.Videos[:0]
	v.PlayList = v.PlayList[:0]
	v.CurPlayListId = ""
	poolPlayListNavRbo.Put(v)
}
