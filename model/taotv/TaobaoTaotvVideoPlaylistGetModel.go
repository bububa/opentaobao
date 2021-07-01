package taotv

// TaobaoTaotvVideoPlaylistGetModel 结构体
type TaobaoTaotvVideoPlaylistGetModel struct {
	// 视频图片
	PicUrl string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
	// 视频ID
	VideoId string `json:"video_id,omitempty" xml:"video_id,omitempty"`
	// 视频来源
	From int64 `json:"from,omitempty" xml:"from,omitempty"`
	// 视频标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 时长，单位秒
	Seconds string `json:"seconds,omitempty" xml:"seconds,omitempty"`
	// ott测更新时间
	OttUpdateTime string `json:"ott_update_time,omitempty" xml:"ott_update_time,omitempty"`
}
