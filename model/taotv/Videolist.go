package taotv

import (
	"sync"
)

// Videolist 结构体
type Videolist struct {
	// 视频ID信息
	VideoId string `json:"video_id,omitempty" xml:"video_id,omitempty"`
	// 视频名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 视频图片
	Pic string `json:"pic,omitempty" xml:"pic,omitempty"`
	// 视频的来源类型，来源这个视频所在节目单的视频来源类型
	VideoExtType int64 `json:"video_ext_type,omitempty" xml:"video_ext_type,omitempty"`
	// 当前视频的播单ID
	PlayListId int64 `json:"play_list_id,omitempty" xml:"play_list_id,omitempty"`
	// 当前视频的节目ID
	ProgramId int64 `json:"program_id,omitempty" xml:"program_id,omitempty"`
	// 当前视频的排序
	Sort int64 `json:"sort,omitempty" xml:"sort,omitempty"`
	// 视频时长（单位秒）
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 主键ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolVideolist = sync.Pool{
	New: func() any {
		return new(Videolist)
	},
}

// GetVideolist() 从对象池中获取Videolist
func GetVideolist() *Videolist {
	return poolVideolist.Get().(*Videolist)
}

// ReleaseVideolist 释放Videolist
func ReleaseVideolist(v *Videolist) {
	v.VideoId = ""
	v.Name = ""
	v.Pic = ""
	v.VideoExtType = 0
	v.PlayListId = 0
	v.ProgramId = 0
	v.Sort = 0
	v.Duration = 0
	v.Id = 0
	poolVideolist.Put(v)
}
