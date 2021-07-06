package media

// VideoItemExtDo 结构体
type VideoItemExtDo struct {
	// 视频播放地址
	PlayUrl string `json:"play_url,omitempty" xml:"play_url,omitempty"`
	// 视频封面-主图
	MainPicUrl string `json:"main_pic_url,omitempty" xml:"main_pic_url,omitempty"`
	// 视频状态
	State int64 `json:"state,omitempty" xml:"state,omitempty"`
	// 视频基本信息
	VideoInfo *VideoItemDo `json:"video_info,omitempty" xml:"video_info,omitempty"`
	// 是否能在移动端播放
	CanPlayInPhone bool `json:"can_play_in_phone,omitempty" xml:"can_play_in_phone,omitempty"`
}
