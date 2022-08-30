package tmallnr

// NrtCrmLiveDto 结构体
type NrtCrmLiveDto struct {
	// 预考链接
	LiveForeShow string `json:"live_fore_show,omitempty" xml:"live_fore_show,omitempty"`
	// 淘口令
	LiveTaoPwd string `json:"live_tao_pwd,omitempty" xml:"live_tao_pwd,omitempty"`
	// 直播间链接
	LiveUrl string `json:"live_url,omitempty" xml:"live_url,omitempty"`
	// 直播开始时间
	LiveEndTime string `json:"live_end_time,omitempty" xml:"live_end_time,omitempty"`
	// 直播结束时间
	LiveStartTime string `json:"live_start_time,omitempty" xml:"live_start_time,omitempty"`
	// 直播间标题
	LiveTitle string `json:"live_title,omitempty" xml:"live_title,omitempty"`
	// 直播间封面URL
	LiveCoverUrl string `json:"live_cover_url,omitempty" xml:"live_cover_url,omitempty"`
	// 直播间状态
	LiveStatus int64 `json:"live_status,omitempty" xml:"live_status,omitempty"`
	// 直播间ID
	LiveId int64 `json:"live_id,omitempty" xml:"live_id,omitempty"`
}
